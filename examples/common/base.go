package common

import (
	"github.com/houzhongjian/work"
	"log"
)


func RecordLog(ctx *work.Context) {
	log.Println("record log")
	ctx.Done()
}

func CheckLogin(ctx *work.Context) {
	log.Println("check login")

	//log.Println("禁止登录")
	//ctx.Done()
}