package v1

import (
	. "github.com/alimy/mir/v4"
	"github.com/alimy/mir/v4/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	engine.Entry[AfdianCallback]()
}

type AfdianCallback struct {
	Group `mir:"v1"`

	// 爱发电支付回掉
	Callback func(Post, web.AfdianCallbackReq) web.AfdianCallbackResp `mir:"/afdian/callback"`
}
