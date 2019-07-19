# qq-miniprogram
[![Build Status](https://travis-ci.org/solarhell/mina.svg?branch=master)](https://travis-ci.org/solarhell/mina)
字节跳动小程序(头条小程序) golang sdk

## done
登录
AccessToken(需持久化 防止超过请求限制)

## todo


## usage

### 登录
```go
package main

import (
	TM "github.com/solarhell/tt-miniprogram"
	"net/http"
	"time"
)

func main() {
	c := TM.NewClient(&http.Client{
		Timeout: 30 * time.Second,
		Transport: &TM.DebugRequestTransport{
			RequestHeader:  true,
			RequestBody:    true,
			ResponseHeader: true,
			ResponseBody:   true,
			Transport: &http.Transport{
				IdleConnTimeout: 30 * time.Second,
	        },
		},
	})

	ui, err := c.Login("appid", "secret", "code")
	...
}
```
