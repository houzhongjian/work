package main

import (
	"github.com/houzhongjian/work"
	"github.com/houzhongjian/work/examples/service"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	web := work.New()
	web.Any("/", &service.DefaultRequest{})
	web.Post("/api/login", &service.LoginRequest{})

	panic(web.Run(":8100"))
}
