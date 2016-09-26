package log4go

import (
	"fmt"
	"io"
	"log"
	"sync"
)

type Log4go struct {
	filterWriters map[level]*log.Logger
	filterReaders map[level][]io.WriteCloser
	store         chan *message
	merge         bool
	wg            sync.WaitGroup
}

func NewLog4go() *Log4go {
	return &Log4go{}
}

// 工作协程
func (this *Log4go) writing() {
	defer this.wg.Done()
	for {
		msg, ok := <-this.store
		if !ok {
			return
		}
		// 降级判断
		for lvl := DEBUG; lvl <= ERROR; lvl++ {
			if lvl > msg.Level {
				break
			}
			if lvl != msg.Level && !this.merge {
				continue
			}
			// 写入缓存
			filters, _ := this.filterReaders[lvl]
			for _, filter := range filters {
				nlen, err := filter.Write(msg.Data)
				if err != nil {
					panic(fmt.Errorf("filter.Write(%s) error(%v)", msg.Data, err))
				}
				if alen := len(msg.Data); nlen != alen {
					panic(fmt.Errorf("nlen(%d) != len(msg.Data)(%d)", nlen, alen))
				}
			}
		}
	}
}

// 关闭
func (this *Log4go) Close() {
	if this.store == nil {
		return
	}
	close(this.store)
	this.wg.Wait()
	this.filterWriters = nil
	for _, filters := range this.filterReaders {
		for _, filter := range filters {
			filter.Close()
		}
	}
	this.filterReaders = nil
	this.store = nil
	this.merge = false
}

func (this *Log4go) Logger(lvl level) *log.Logger {
	if filter, ok := this.filterWriters[lvl]; ok {
		return filter
	}
	return nil
}

func (this *Log4go) Error(format string, v ...interface{}) {
	logger := this.Logger(ERROR)
	if logger == nil {
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}

func (this *Log4go) Warn(format string, v ...interface{}) {
	logger := this.Logger(WARNING)
	if logger == nil {
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}

func (this *Log4go) Info(format string, v ...interface{}) {
	logger := this.Logger(INFO)
	if logger == nil {
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}

func (this *Log4go) Trace(format string, v ...interface{}) {
	logger := this.Logger(TRACE)
	if logger == nil {
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}

func (this *Log4go) Debug(format string, v ...interface{}) {
	logger := this.Logger(DEBUG)
	if logger == nil {
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}
