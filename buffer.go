package log4go

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"text/template"
	"time"
)

// 写入缓冲区
type Buffer struct {
	tmpl  *template.Template
	level level
	store chan *message
}

// 新建写入缓冲区
func NewBuffer(format string, level level, store chan *message) *Buffer {
	// 加载适配器
	tmplText := adapter.Replace(string(format))
	if len(format) == 0 || format[len(format)-1] != '\n' {
		tmplText += "\n"
	}
	// 生成适配器模版
	tmpl, err := template.New("Log").Parse(tmplText)
	if err != nil {
		panic(err)
	}
	return &Buffer{tmpl, level, store}
}

// 写入数据
func (this *Buffer) Write(p []byte) (n int, err error) {
	defer func() {
		if e := recover(); e != nil {
			strErr, _ := e.(string)
			n = 0
			err = errors.New(strErr)
		}
	}()
	util := make(map[string]interface{})
	util["LEVEL"] = levelStrings[this.level]
	util["TIME"] = time.Now()
	util["MESSAGE"] = string(p[:len(p)-1])
	// 获取代码所在文件名及行数
	if _, file, line, ok := runtime.Caller(3); ok {
		util["SOURCE"] = fmt.Sprintf("%s:%d", file, line)
	}
	// 数据写入缓存
	buf := bytes.NewBufferString("")
	this.tmpl.Execute(buf, util)
	// 写入数据中心队列
	this.store <- &message{this.level, buf.Bytes()}
	return len(p), nil
}
