package work

type HandlerFunc func(ctx *Context)

var dist = make(map[string]HandlerFunc)

func (*Engine) Method(path string, handler HandlerFunc) {
	dist[path] = handler
}
