package log4go

import (
	"os"
)

// 控制台
type logConsole struct {
	file *os.File
}

func newLogConsole() *logConsole {
	return &logConsole{os.Stdout}
}

// 写入控制台
func (this *logConsole) Write(p []byte) (n int, err error) {
	return this.file.Write(p)
}

// 关闭
func (this *logConsole) Close() error {
	this.file = nil
	return nil
}
