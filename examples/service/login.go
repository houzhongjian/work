package service

import (
	"errors"
	"github.com/houzhongjian/work"
	"github.com/houzhongjian/work/examples/basic"
	"log"
	"net/http"
)

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (request *LoginRequest) Validator() error {
	if request.Account == "" {
		return errors.New("账号不能为空")
	}

	if request.Password == "" {
		return errors.New("密码不能为空")
	}

	return nil
}

func (request *LoginRequest) Before(ctx *work.Context) {
	basic.Common(ctx)
	if err := ctx.BindJSONAndValidator(request); err != nil {
		log.Printf("err:%+v\n", err)

		ctx.WriteHeader(http.StatusForbidden)
		ctx.ServeJSON(work.Message{
			"code":    1001,
			"message": err.Error(),
		})
		ctx.Done()
		return
	}
}

func (request *LoginRequest) Logic(ctx *work.Context) {
	if request.Account == "" {
		ctx.ServeJSON(work.Message{
			"code":    1001,
			"message": "登录失败",
		})
		ctx.Done()
		return
	}

	ctx.ServeJSON(work.Message{
		"code":    1000,
		"message": "登录成功",
	})
}

func (request *LoginRequest) After(ctx *work.Context) {
	log.Printf("After request:%+v\n", request)
}
