package tt_miniprogram

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
)

type watermark struct {
	AppID     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

// Userinfo 解密后的用户信息
type Userinfo struct {
	Openid    string    `json:"openId"`
	Avatar    string    `json:"avatarUrl"`
	Nickname  string    `json:"nickName"`
	Gender    int       `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	Language  string    `json:"language"`
	Watermark watermark `json:"watermark"`
}

// https://developer.toutiao.com/docs/open/dataCodec.html#%E6%A0%A1%E9%AA%8C%E6%95%B0%E6%8D%AE%E5%90%88%E6%B3%95%E6%80%A7
func validate(rawData, sessionKey, signature string) bool {
	r := sha1.Sum([]byte(rawData + sessionKey))

	return signature == hex.EncodeToString(r[:])
}

// https://developer.toutiao.com/docs/open/dataCodec.html#%E8%A7%A3%E5%AF%86%E6%95%8F%E6%84%9F%E6%95%B0%E6%8D%AE
func DecryptUserInfo(rawData, encryptedData, signature, iv, sessionKey string) (ui Userinfo, err error) {
	if ok := validate(rawData, sessionKey, signature); !ok {
		err = errors.New("数据校验失败")
		return
	}

	bts, err := CBCDecrypt(sessionKey, encryptedData, iv)
	if err != nil {
		return
	}

	err = json.Unmarshal(bts, &ui)
	return
}

func CBCDecrypt(ssk, data, iv string) (bts []byte, err error) {
	key, err := base64.StdEncoding.DecodeString(ssk)
	if err != nil {
		return
	}

	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}

	rawIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	size := aes.BlockSize

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < size {
		err = errors.New("cipher too short")
		return
	}

	// CBC mode always works in whole blocks.
	if len(ciphertext)%size != 0 {
		err = errors.New("cipher is not a multiple of the block size")
		return
	}

	mode := cipher.NewCBCDecrypter(block, rawIV[:size])
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	return PKCS5UnPadding(plaintext)
}

// PKCS5UnPadding 反补
// Golang AES没有64位的块, 如果采用PKCS5, 那么实质上就是采用PKCS7
func PKCS5UnPadding(plaintext []byte) ([]byte, error) {
	ln := len(plaintext)

	// 去掉最后一个字节 unPadding 次
	unPadding := int(plaintext[ln-1])

	if unPadding > ln {
		return []byte{}, errors.New("数据不正确")
	}

	return plaintext[:(ln - unPadding)], nil
}
