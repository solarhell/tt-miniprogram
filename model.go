package tt_miniprogram

// https://developer.toutiao.com/docs/server/

// Response 基础数据
type CommonResponse struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type LoginResponse struct {
	CommonResponse
	SessionKey      string `json:"session_key"`
	Openid          string `json:"openid"`
	AnonymousOpenid string `json:"anonymous_openid,omitempty"`
}

type AccessToken struct {
	CommonResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
