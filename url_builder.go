package tt_miniprogram

import "net/url"

const (
	// BaseURL 基础URL
	baseURL  = "https://developer.toutiao.com"
	codeAPI  = "/api/apps/jscode2session"
	tokenAPI = "/api/apps/token"
)

func CodeToURL(appId, appSecret, code, anonymousCode string) (s string, err error) {
	if appId == "" || appSecret == "" || (code == "" && anonymousCode == "") {
		return s, ErrNotAllowEmptyParam
	}
	u, err := url.Parse(baseURL + codeAPI)
	if err != nil {
		return s, err
	}

	query := u.Query()

	query.Set("appid", appId)
	query.Set("secret", appSecret)
	if code != "" {
		query.Set("code", code)
	}
	if anonymousCode != "" {
		query.Set("anonymous_code", anonymousCode)
	}

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func TokenURL(appId, appSecret string) (s string, err error) {
	if appId == "" || appSecret == "" {
		return s, ErrNotAllowEmptyParam
	}
	u, err := url.Parse(baseURL + tokenAPI)
	if err != nil {
		return s, err
	}

	query := u.Query()

	query.Set("appid", appId)
	query.Set("secret", appSecret)
	query.Set("grant_type", "client_credential")

	u.RawQuery = query.Encode()

	return u.String(), nil
}
