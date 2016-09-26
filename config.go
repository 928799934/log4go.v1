// Copyright (C) 2010, Kyle Lemons <kyle@kylelemons.net>.  All rights reserved.

package log4go

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	ErrAttr  = "LoadConfiguration: Error: Required attribute %s for filter missing in %s\n"
	ErrFile  = "LoadConfiguration: Error: Could not parse XML configuration in %q: %s\n"
	ErrChild = "LoadConfiguration: Error: Required child <%s> for filter missing in %s\n"
)

// 加载配置
func (this *Log4go) LoadConfiguration(filename string) error {
	this.Close()
	// Open the configuration file
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	xc := new(xmlLoggerConfig)
	if err := xml.Unmarshal(contents, xc); err != nil {
		return fmt.Errorf(ErrFile, filename, err)
	}

	filterWriters := make(map[level]*log.Logger)
	filterReaders := make(map[level][]io.WriteCloser)
	store := make(chan *message, 10240)
	var (
		lvl             level
		format          string
		objWriterCloser io.WriteCloser
	)

	for _, xmlfilt := range xc.Filter {
		// 判断过滤器是否可用
		if !xmlfilt.Enabled {
			continue
		}

		// 识别等级
		switch xmlfilt.Level {
		case "DEBUG":
			lvl = DEBUG
		case "TRACE":
			lvl = TRACE
		case "INFO":
			lvl = INFO
		case "WARNING":
			lvl = WARNING
		case "ERROR":
			lvl = ERROR
		default:
			return fmt.Errorf(ErrChild, "level", filename)
		}

		// 判断支持的输出类型
		switch xmlfilt.Type {
		case "console":
			format, objWriterCloser = xmlToConsoleLogWriter(xmlfilt.Property)
		case "file":
			format, objWriterCloser = xmlToFileLogWriter(xmlfilt.Property)
		default:
			return fmt.Errorf(ErrChild, "type", filename)
		}

		filterWriters[lvl] = log.New(NewBuffer(format, lvl, store), "", 0)
		filterReaders[lvl] = append(filterReaders[lvl], objWriterCloser)
	}
	this.filterWriters = filterWriters
	this.filterReaders = filterReaders
	this.store = store
	this.merge = !xc.Unmerge
	this.wg.Add(1)
	go this.writing()
	return nil
}

// Parse a number with K/M/G suffixes based on thousands (1000) or 2^10 (1024)
// Parse a number with S/M/H suffixes based on thousands (60)
func strToNumSuffix(str string, mult int) int {
	num := 1
	if len(str) < 2 {
		num *= mult
		parsed, _ := strconv.Atoi(str)
		return parsed * num
	}

	switch str[len(str)-1] {
	case 'G', 'g', 'H', 'h':
		num *= mult
		fallthrough
	case 'M', 'm':
		num *= mult
		fallthrough
	case 'K', 'k', 'S', 's':
		num *= mult
		str = str[0 : len(str)-1]
	}
	parsed, _ := strconv.Atoi(str)
	return parsed * num
}

func xmlToFileLogWriter(props []xmlProperty) (string, io.WriteCloser) {
	file := ""
	format := "[%D %T] [%L] (%S) %M"
	maxsize := 0
	delay := 0
	// Parse properties
	for _, prop := range props {
		switch prop.Name {
		case "filename":
			file = strings.Trim(prop.Value, " \r\n")
		case "format":
			format = strings.Trim(prop.Value, " \r\n")
		case "maxsize":
			maxsize = strToNumSuffix(strings.Trim(prop.Value, " \r\n"), 1024)
		case "delay":
			delay = strToNumSuffix(strings.Trim(prop.Value, " \r\n"), 60)
		}
	}
	// Check properties
	if len(file) == 0 {
		panic(fmt.Errorf("filetype: filename is empty"))
	}
	return format, NewFile(file, maxsize, delay)
}

func xmlToConsoleLogWriter(props []xmlProperty) (string, io.WriteCloser) {
	format := "[%D %T] [%L] (%S) %M"
	// Parse properties
	for _, prop := range props {
		switch prop.Name {
		case "format":
			format = strings.Trim(prop.Value, " \r\n")
		}
	}
	return format, NewConsole()
}
