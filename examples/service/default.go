package service

import (
	"github.com/houzhongjian/work"
	"log"
)

type DefaultRequest struct {}

func HandleDefault(ctx *work.Context)  {
	srv := new(DefaultRequest)
	ctx.Step(srv.Before, srv.Logic, srv.After)
}

func (request *DefaultRequest) Before(ctx *work.Context) {
	log.Printf("Before request:%+v\n", request)
}

func (request *DefaultRequest) After(ctx *work.Context) {
	log.Printf("After request:%+v\n", request)
}

func (request *DefaultRequest) Logic(ctx *work.Context) {
	ctx.ResponseWriter.Write([]byte("<h1>hello world</h1>"))
}
