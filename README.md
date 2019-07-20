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
)

func main() {

	ui, err := TM.Login("appid", "secret", "code")
	...
}
```
