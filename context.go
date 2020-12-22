package work

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	next           bool
}

//WriteResponse 返回自定义消息.
func (ctx *Context) WriteResponse(code int, msg interface{}) {
	data := map[string]interface{}{
		"code": code,
		"msg":  msg,
	}

	content, err := json.Marshal(data)
	if err != nil {
		return
	}

	_, _ = ctx.ResponseWriter.Write(content)
}

//WriteFail 用来返回失败消息 并且阻止执行下一个步骤方法.
func (ctx *Context) WriteFail(msg interface{}) {
	ctx.Done()
	ctx.WriteResponse(1001, msg)
}

//WriteWarning 用来返回失败消息.
func (ctx *Context) WriteWarning(msg interface{}) {
	ctx.WriteResponse(1001, msg)
}

//返回成功消息.
func (ctx *Context) WriteSuccess(msg interface{}) {
	ctx.WriteResponse(1000, msg)
}

func (ctx *Context) BindJSON(obj interface{}) error {
	if err := json.NewDecoder(ctx.Request.Body).Decode(&obj); err != nil {
		return err
	}
	return nil
}

//BindJSONAndValidator 绑定json 并且调用结构体的Validator.
func (ctx *Context) BindJSONAndValidator(obj interface{}) error {
	if err := json.NewDecoder(ctx.Request.Body).Decode(obj); err != nil {
		return err
	}

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	fn := "Validator"
	_, ok := t.MethodByName(fn)
	if !ok {
		//不存在.
		return errors.New("Validator不存在")
	}

	validator := v.MethodByName(fn)
	val := validator.Call(nil) //返回Value类型

	if len(val) > 0 && val[0].Interface() != nil {
		return val[0].Interface().(error)
	}
	return nil
}

func (ctx *Context) Step(args ...HandlerFunc) {
	for _, arg := range args {
		if ctx.next { //判断是否运行执行下一步.
			arg(ctx)
		}
	}
}

func (ctx *Context) Done(args ...HandlerFunc) {
	ctx.next = false
}
