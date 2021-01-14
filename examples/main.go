package main

import (
	"github.com/houzhongjian/work"
	"github.com/houzhongjian/work/examples/service"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	web := work.New()
	web.Post("/api/login", &service.LoginRequest{})
	web.Get("/", &service.DefaultRequest{})
	web.Delete("/api/admin/delete", &service.DefaultRequest{})
	web.Post("/api/admin/add", &service.DefaultRequest{})
	web.Get("/api/admin/list", &service.DefaultRequest{})
	web.NotFound(&service.NotFoundRequest{})

	panic(web.Run(":8100"))
}
