package work

import (
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Engine struct {
	router   map[string]Route
	notfound interface{}
	filter []HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: map[string]Route{
			http.MethodGet:    getRouter,
			http.MethodPut:    putRouter,
			http.MethodPost:   postRouter,
			http.MethodDelete: deleteRouter,
		},
	}
}

func (engine *Engine) Filter(filter ...HandlerFunc) {
	engine.filter = filter
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		ResponseWriter: w,
		Request:        r,
		next:           true,
		Body:           r.Body,
		contextData:    make(map[string]interface{}),
		httpStatus:     http.StatusOK,
		RequestURI:     r.RequestURI,
		Method:         r.Method,
		Header:         r.Header,
		Host:           r.Host,
	}

	//执行过滤器.
	for _, filter := range engine.filter {
		if ctx.next {
			filter(ctx)
		}
	}
	if !ctx.next {
		return
	}

	urlPath := strings.Split(r.URL.String(), "?")

	//判断method是否存在.
	method, ok := engine.router[r.Method]
	if !ok {
		//todo 执行404
		http.NotFound(w, r)
		return
	}

	//判断路由是否存在.
	obj, ok := method[urlPath[0]]
	if !ok {
		http.NotFound(w, r)
		return
	}

	t := reflect.TypeOf(obj)
	//判断结构体是否为指针.
	if t.Kind() != reflect.Ptr {
		log.Println(t.Name(), "必须为指针类型")
		return
	}
	v := reflect.New(reflect.ValueOf(obj).Elem().Type())

	_, ok = t.MethodByName("Before")
	if ok {
		args := []reflect.Value{reflect.ValueOf(ctx)}
		value := v.MethodByName("Before")
		value.Call(args)
	}

	if ctx.next {
		_, ok = t.MethodByName("Logic")
		if ok {
			args := []reflect.Value{reflect.ValueOf(ctx)}
			value := v.MethodByName("Logic")
			value.Call(args)
		}
	}

	if ctx.next {
		_, ok = t.MethodByName("After")
		if ok {
			args := []reflect.Value{reflect.ValueOf(ctx)}
			value := v.MethodByName("After")
			value.Call(args)
		}
	}
}
