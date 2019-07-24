package tt_miniprogram

import "testing"

func TestDecryptUserInfo(t *testing.T) {
	code := "xxxx" // Login 一次就会失效

	encryptedData := "xxx=="
	iv := "xxx=="
	rawData := `{"nickName":"xxxx","avatarUrl":"xxxx","gender":0,"city":"","province":"","country":"中国","language":""}`
	signature := "c838a00593e7b0c51acf956ed89e17ab4af2b89a"

	appId := "ttbxxxxxxx"
	appSecret := "xxxxxx"

	lr, err := Login(appId, appSecret, code, "")
	if err != nil {
		t.Error(err)
	}

	ui, err := DecryptUserInfo(rawData, encryptedData, signature, iv, lr.SessionKey)
	if err != nil {
		t.Error(err)
	}

	t.Log("Openid", ui.Openid)
}
