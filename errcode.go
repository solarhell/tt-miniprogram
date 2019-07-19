package tt_miniprogram

import "errors"

var (
	ErrNotAllowEmptyParam     = errors.New("param cannot be empty")
	ErrConnectByteDanceServer = errors.New("err connecet ByteDance server")
)
