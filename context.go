package work

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strconv"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	next           bool //next用来中断调用链.
	Body           io.ReadCloser
	contextData    map[string]interface{}
	httpStatus     int
	RequestURI     string
	RemoteAddr     string
	Method         string
	Header         http.Header
	Host           string
}

func (ctx *Context) ServeJSON(h H) {
	b, err := json.Marshal(h)
	if err != nil {
		return
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.Write(b)
}

func (ctx *Context) ServeData(h H) {
}

func (ctx *Context) ServeString(s string) {
	ctx.Write([]byte(s))
}

func (ctx *Context) WriteHeader(statusCode int) {
	ctx.ResponseWriter.WriteHeader(statusCode)
}

func (ctx *Context) BindJSON(obj interface{}) error {
	if err := json.NewDecoder(ctx.Request.Body).Decode(obj); err != nil {
		return err
	}
	return nil
}

//BindJSONAndValidator 绑定json 并且调用结构体的Validator.
func (ctx *Context) BindJSONAndValidator(obj interface{}) error {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		return errors.New(t.Name() + "必须为指针类型")
	}
	v := reflect.ValueOf(obj)

	fn := "Validator"
	_, ok := t.MethodByName(fn)
	if !ok {
		//不存在.
		return errors.New("Validator不存在")
	}

	if err := json.NewDecoder(ctx.Request.Body).Decode(obj); err != nil {
		return err
	}

	validator := v.MethodByName(fn)
	val := validator.Call(nil) //返回Value类型

	if len(val) > 0 && val[0].Interface() != nil {
		return val[0].Interface().(error)
	}
	return nil
}

//Done 方法 主要用来中断Step中的方法 不在往下执行.
func (ctx *Context) Done() {
	ctx.next = false
}

func (ctx *Context) Write(b []byte) {
	_, _ = ctx.ResponseWriter.Write(b)
}

func (ctx *Context) GetString(key string) string {
	return ctx.Request.FormValue(key)
}

func (ctx *Context) GetInt(key string) int {
	n, err := strconv.Atoi(ctx.Request.FormValue(key))
	if err != nil {
		return 0
	}

	return n
}

func (ctx *Context) SetContext(key string, value interface{}) {
	ctx.contextData[key] = value
}
func (ctx *Context) GetContext(key string) interface{} {
	return ctx.contextData[key]
}
