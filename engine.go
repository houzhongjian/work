package work

import (
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Engine struct {
	router map[string]interface{}
}

func New() *Engine {
	return &Engine{
		router: make(map[string]interface{}),
	}
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := strings.Split(r.URL.String(), "?")
	obj, ok := engine.router[urlPath[0]]
	if !ok {
		http.NotFound(w, r)
		return
	}

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
