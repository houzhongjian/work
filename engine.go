package work

import "net/http"

type Engine struct {
}

func New() *Engine {
	return new(Engine)
}

func (w *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, w)
}

func (*Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := dist[r.URL.String()]
	if !ok {
		return
	}

	f(&Context{
		ResponseWriter: w,
		Request:        r,
		next:           true,
	})
}
