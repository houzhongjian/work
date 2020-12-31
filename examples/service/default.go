package service

import (
	"github.com/houzhongjian/work"
	"github.com/houzhongjian/work/examples/basic"
)

type DefaultRequest struct {
	basic.PageParams
}

func (request *DefaultRequest) Before(ctx *work.Context) {
	request.LoadPageParams(ctx)
}

func (request *DefaultRequest) After(ctx *work.Context) {
}

func (request *DefaultRequest) Logic(ctx *work.Context) {
	ctx.ServeJSON(work.H{
		"page":     request.Page,
		"pagesize": request.Pagesize,
		"offset":   request.Offset,
	})
}
