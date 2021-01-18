package common

import (
	"github.com/houzhongjian/work"
	"log"
)

func RecordLog(ctx *work.Context) {
	log.Println("record log")
	//ctx.Done()
}

func CheckLogin(ctx *work.Context) {
	log.Println("check login")

	//log.Println("禁止登录")
	//ctx.Done()
	if ctx.RequestURI== "/api/login" || ctx.RequestURI == "/" {
		return
	}

	token := ctx.Header.Get("token")
	if token == "" {
		ctx.ServeJSON(work.Message{
			"code":    1001,
			"message": "未登录",
		})
		ctx.Done()
	}
}
