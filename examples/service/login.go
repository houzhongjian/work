package service

import (
	"github.com/houzhongjian/work"
	"log"
	"net/http"
)

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (*LoginRequest) Validator() error {
	return nil
}

func (request *LoginRequest) Before(ctx *work.Context) {
	if err := ctx.BindJSONAndValidator(request); err != nil {
		log.Printf("err:%+v\n", err)

		ctx.WriteHeader(http.StatusForbidden)
		ctx.ServeJSON(work.H{
			"code":    1001,
			"message": err.Error(),
		})
		ctx.Done()
		return
	}
}

func (request *LoginRequest) Logic(ctx *work.Context) {
	if request.Account == "" {
		ctx.ServeJSON(work.H{
			"code":    1001,
			"message": "登录失败",
		})
		ctx.Done()
		return
	}

	ctx.ServeJSON(work.H{
		"code":    1000,
		"message": "登录成功",
	})
}

func (request *LoginRequest) After(ctx *work.Context) {
	log.Printf("After request:%+v\n", request)
}
