package basic

import (
	"github.com/houzhongjian/work"
)

type PageParams struct {
	Page     int
	Pagesize int
	Offset   int
}

func (request *PageParams) LoadPageParams(ctx *work.Context) {
	request.Page = ctx.GetInt("page")
	request.Pagesize = ctx.GetInt("pagesize")

	if request.Page < 1 {
		request.Page = 1
	}

	if request.Pagesize < 1 {
		request.Pagesize = 20
	}

	request.Offset = (request.Page - 1) * request.Pagesize
}
