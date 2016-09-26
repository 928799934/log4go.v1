package log4go

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// 文件
type File struct {
	ext        string
	name       string
	path       string
	size       int
	delay      int
	day        string
	num        int
	maxsize    int
	file       *os.File
	wg         sync.WaitGroup
	closeEvent chan bool
}

func NewFile(name string, size, delay int) *File {
	// 取目录
	dir := filepath.Dir(name)
	// 取后缀名
	ext := filepath.Ext(name)
	// 取基础路径
	base := filepath.Base(name)
	// 取文件名
	name = strings.TrimSuffix(base, ext)
	f := &File{
		path:       dir,
		ext:        ext,
		name:       name,
		maxsize:    size,
		delay:      delay,
		closeEvent: make(chan bool, 1),
	}
	// 打开文件
	if err := f.Open(); err != nil {
		panic(err)
	}

	// 是否延迟
	if f.delay > 0 {
		f.wg.Add(1)
		go f.flush()
	}
	return f
}

// 文件大小 控制
func (this *File) fixSize(nlen int) error {
	if this.maxsize == 0 || this.size+nlen < this.maxsize {
		return nil
	}
	old := this.file.Name()
	new := fmt.Sprintf("%s/%s-%s-%d%s", this.path, this.name, this.day, this.num,
		this.ext)
	this.num++
	this.file.Close()
	this.file = nil
	this.size = 0
	return os.Rename(old, new)
}

// 打开文件
func (this *File) Open() (err error) {
	now := time.Now().Format("2006-01-02")
	if now != this.day {
		if this.file != nil {
			this.file.Close()
			this.file = nil
		}
		this.day = now
		this.size = 0
		this.num = 0
	}
	if this.file != nil {
		return
	}
	name := fmt.Sprintf("%s/%s-%s%s", this.path, this.name, this.day, this.ext)
	this.file, err = os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	return
}

// 刷文件到磁盘
func (this *File) flush() {
	defer this.wg.Done()
	delay := time.Second * time.Duration(this.delay)
	t := time.NewTimer(delay)
	for {
		select {
		case <-this.closeEvent:
			this.file.Sync()
			return
		case <-t.C:
			this.file.Sync()
			t.Reset(delay)
		}
	}
}

// 写入
func (this *File) Write(p []byte) (n int, err error) {
	this.Open()
	this.fixSize(len(p))
	n, err = this.file.Write(p)
	if err != nil {
		return
	}
	this.size += n
	if this.delay == 0 {
		this.file.Sync()
	}
	return
}

// 关闭
func (this *File) Close() error {
	close(this.closeEvent)
	this.wg.Wait()
	this.file.Close()
	this.file = nil
	return nil
}
