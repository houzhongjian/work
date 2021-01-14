package work

import (
	"net/http"
)

type Route map[string]interface{}

var (
	getRouter    = make(Route)
	postRouter   = make(Route)
	deleteRouter = make(Route)
	putRouter    = make(Route)
)

type HandlerFunc func(ctx *Context)

func (engine *Engine) Delete(path string, handler interface{}) {
	deleteRouter := engine.router[http.MethodDelete]
	deleteRouter[path] = handler
}

func (engine *Engine) Get(path string, handler interface{}) {
	getRouter := engine.router[http.MethodGet]
	getRouter[path] = handler
}

func (engine *Engine) Post(path string, handler interface{}) {
	postRouter := engine.router[http.MethodPost]
	postRouter[path] = handler
}

func (engine *Engine) Put(path string, handler interface{}) {
	putRouter := engine.router[http.MethodPut]
	putRouter[path] = handler
}

func (engine *Engine) NotFound(handler interface{}) {
	engine.notfound = handler
}
