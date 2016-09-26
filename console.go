package log4go

import (
	"os"
)

// 控制台
type Console struct {
	file *os.File
}

func NewConsole() *Console {
	return &Console{os.Stdout}
}

// 写入控制台
func (this *Console) Write(p []byte) (n int, err error) {
	return this.file.Write(p)
}

// 关闭
func (this *Console) Close() error {
	this.file = nil
	return nil
}
