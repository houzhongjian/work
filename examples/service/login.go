package service

import (
	"github.com/houzhongjian/work"
	"log"
)

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func HandleLogin(ctx *work.Context) {
	srv := new(LoginRequest)
	ctx.Step(srv.Before, srv.Logic, srv.After)
}

func (request *LoginRequest) Before(ctx *work.Context) {
	log.Printf("Before request:%+v\n", request)
	if err := ctx.BindJSON(request); err != nil {
		log.Printf("err:%+v\n", err)
		ctx.WriteFail(err.Error())
		return
	}
}

func (request *LoginRequest) After(ctx *work.Context) {
	log.Printf("After request:%+v\n", request)
}

func (request *LoginRequest) Logic(ctx *work.Context) {
	request.Account = "zhangsan"
	request.Password = "123456"
	log.Printf("Logic request:%+v\n", request)
}
