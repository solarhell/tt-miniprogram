package tt_miniprogram

import (
	"encoding/json"
	"errors"
)

func (c *Client) Login(appId, appSecret, code, anonymousCode string) (lr LoginResponse, err error) {
	api, err := CodeToURL(appId, appSecret, code, anonymousCode)
	if err != nil {
		return lr, err
	}

	res, err := c.client.Get(api)

	if err != nil {
		return lr, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return lr, ErrConnectByteDanceServer
	}

	err = json.NewDecoder(res.Body).Decode(&lr)
	if err != nil {
		return lr, err
	}

	if lr.Errcode != 0 {
		return lr, errors.New(lr.Errmsg)
	}

	return lr, nil
}
