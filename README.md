# Server Push [![GoDoc](https://godoc.org/github.com/go-mego/push?status.svg)](https://godoc.org/github.com/go-mego/push)

Server Push 是基於 HTTP/2 協定的檔案主動推送，這能夠讓你主動將客戶端可能會載入的靜態檔案主動推送給他們，而不是等到需要時客戶端才會要求。

# 索引

* [安裝方式](#安裝方式)
* [使用方式](#使用方式)

# 安裝方式

打開終端機並且透過 `go get` 安裝此套件即可。

```bash
$ go get github.com/go-mego/push
```

# 使用方式

先透過 `push.New()` 初始化一個 Server Push 推送中介軟體，並將其傳入 Mego 中的 `Use()` 函式來表明要套用到該引擎中供稍後的路由內使用。以 `Push` 並指定靜態檔案來開始推送。

```go
package main

import (
	"github.com/go-mego/mego"
	"github.com/go-mego/push"
)

func main() {
    m := mego.New()
    // 將 Server Push 套用到全域中介軟體裡供稍後在路由中使用。
    m.Use(push.New())
	m.Get("/", func(p *push.Pusher) string {
		// 透過 `Push` 來主動向客戶端推送指定的靜態檔案。
		if err := p.Push("/example.png"); err != nil {
			return "Push 檔案時發生錯誤！"
		}
		if err := p.Push("/favorite.mp3"); err != nil {
			return "Push 檔案時發生錯誤！"
		}
		return "檔案都已經成功地 Push 了！"
	})
	m.Run()
}
```