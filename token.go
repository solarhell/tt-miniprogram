package tt_miniprogram

import (
	"errors"
	"github.com/imroc/req"
)

func GetAccessToken(appId, appSecret string) (ak AccessToken, err error) {
	api, err := TokenURL(appId, appSecret)
	if err != nil {
		return ak, err
	}

	r, err := req.Get(api)
	if err != nil {
		return ak, err
	}

	if r.Response().StatusCode != 200 {
		return ak, ErrConnectByteDanceServer
	}

	err = r.ToJSON(&ak)
	if err != nil {
		return ak, err
	}

	if ak.Errcode != 0 {
		return ak, errors.New(ak.Errmsg)
	}

	return ak, nil
}
