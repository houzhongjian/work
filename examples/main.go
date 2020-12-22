package main

import (
	"github.com/houzhongjian/work"
	"github.com/houzhongjian/work/examples/service"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	e := work.New()
	e.Method("/", service.HandleDefault)
	e.Method("/login", service.HandleLogin)
	panic(e.Run(":9099"))
}
