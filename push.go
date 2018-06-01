package main

import (
	"net/http"

	"github.com/go-mego/mego"
)

// New 會回傳一個新的 Push 中介軟體。
func New() mego.HandlerFunc {
	return func(c *mego.Context) {
		c.Injector.Map(&Pusher{
			context: c,
		})
		c.Next()
	}
}

// Pusher 重現了一個 Push 中介軟體。
type Pusher struct {
	context *mego.Context
}

// Push 能夠將指定的靜態檔案主動推送至客戶端瀏覽器。
// 當客戶端不支援 HTTP/2 的 Server Push 功能時會回傳 `http.ErrNotSupported` 錯誤。
func (p *Pusher) Push(path string) error {
	return p.context.Writer.Push(path, nil)
}

// PushWithHeader 能夠將指定的靜態檔案與指定的標頭一同主動推送至客戶端瀏覽器。
// 當客戶端不支援 HTTP/2 的 Server Push 功能時會回傳 `http.ErrNotSupported` 錯誤。
func (p *Pusher) PushWithHeader(path string, header map[string][]string) error {
	return p.context.Writer.Push(path, &http.PushOptions{
		Header: header,
	})
}
