package work

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"strings"
)

func (ctx *Context) Layout(name string) {
	ctx.layout = name
}

func (ctx *Context) ServeData(data Message) {
	ctx.viewData = data
}

func (ctx *Context) View(view string) {
	ctx.view = view

	//读取layout.
	layoutByte, err := ioutil.ReadFile(ctx.layout)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}

	//读取视图文件.
	viewByte, err := ioutil.ReadFile(view)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}

	//替换{{.LayoutContent}}.
	content := strings.ReplaceAll(string(layoutByte), "{{.LayoutContent}}", string(viewByte))

	t := template.New("layout.html").Funcs(template.FuncMap{
		"include": include,
	})
	t, err = t.Parse(content)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}

	err = t.Execute(ctx.ResponseWriter, ctx.viewData)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}
}

//template include.
func include(args string, v ...interface{}) template.HTML {
	buf := new(bytes.Buffer)
	t, err := template.ParseFiles(args)
	if err != nil {
		log.Printf("err:%+v\n", err)
	}

	if len(v) > 0 {
		if err := t.Execute(buf, v[0]); err != nil {
			log.Printf("err:%+v\n", err)
		}
	} else {
		if err := t.Execute(buf, nil); err != nil {
			log.Printf("err:%+v\n", err)
		}
	}
	return template.HTML(buf.String())
}
