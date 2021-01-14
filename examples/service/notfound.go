package service

import (
	"github.com/houzhongjian/work"
	"log"
	"net/http"
)

type NotFoundRequest struct{}

func (request *NotFoundRequest) Before(ctx *work.Context) {
}

func (request *NotFoundRequest) Logic(ctx *work.Context) {
	ctx.ServeJSON(work.Message{
		"code":    1001,
		"message": "当前页面不存在",
	})
	ctx.WriteHeader(http.StatusNotFound)
}

func (request *NotFoundRequest) After(ctx *work.Context) {
	log.Printf("After request:%+v\n", request)
}
