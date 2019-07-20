package tt_miniprogram

import (
	"errors"
	"github.com/imroc/req"
)

func Login(appId, appSecret, code, anonymousCode string) (lr LoginResponse, err error) {
	api, err := CodeToURL(appId, appSecret, code, anonymousCode)
	if err != nil {
		return lr, err
	}

	r, err := req.Get(api)
	if err != nil {
		return lr, err
	}

	if r.Response().StatusCode != 200 {
		return lr, ErrConnectByteDanceServer
	}

	err = r.ToJSON(&lr)
	if err != nil {
		return lr, err
	}

	if lr.Errcode != 0 {
		return lr, errors.New(lr.Errmsg)
	}

	return lr, nil
}
