package log4go

import (
	"bytes"
	"errors"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// 写入缓冲区
type logBuffer struct {
	tmpl  *template.Template
	level level
	store chan *message
	opts  int64
}

// 新建写入缓冲区
func newLogBuffer(format string, level level, store chan *message) *logBuffer {
	var opts int64
	for k, v := range options {
		if -1 == strings.Index(format, k) {
			continue
		}
		opts = opts | v
	}
	// 替换适配器
	tmplText := adapter.Replace(format) + "\n"
	// 生成适配器模版
	tmpl, err := template.New("Log").Parse(tmplText)
	if err != nil {
		panic(err)
	}
	return &logBuffer{tmpl, level, store, opts}
}

// 写入数据
func (this *logBuffer) Write(p []byte) (n int, err error) {
	defer func() {
		if e := recover(); e != nil {
			strErr, _ := e.(string)
			n = 0
			err = errors.New(strErr)
		}
	}()
	util := make(map[string]interface{})
	if this.opts&t == t || this.opts&d == d {
		util["TIME"] = time.Now()
	}
	if this.opts&l == l {
		util["LEVEL"] = levelStrings[this.level]
	}
	if this.opts&m == m {
		util["MESSAGE"] = string(p[:len(p)-1])
	}
	if this.opts&s == s {
		// 获取代码所在文件名及行数
		if _, file, line, ok := runtime.Caller(3); ok {
			util["SOURCE"] = file + ":" + strconv.Itoa(line)
		}
	}
	// 数据写入缓存
	buf := bytes.NewBufferString("")
	err = this.tmpl.Execute(buf, util)
	if err != nil {
		panic(err)
	}
	// 写入数据中心队列
	this.store <- &message{this.level, buf.Bytes()}
	return len(p), nil
}
