package work

type HandlerFunc func(ctx *Context)

func (engine *Engine) Any(path string, handler interface{}) {
	engine.router[path] = handler
}

func (engine *Engine) Delete(path string, handler interface{}) {
	engine.router[path] = handler
}

func (engine *Engine) Get(path string, handler interface{}) {
	engine.router[path] = handler
}

func (engine *Engine) Post(path string, handler interface{}) {
	engine.router[path] = handler
}

func (engine *Engine) Put(path string, handler interface{}) {
	engine.router[path] = handler
}
