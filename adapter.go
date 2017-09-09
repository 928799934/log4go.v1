package log4go

import (
	"strings"
)

// template 模版适配器
const (
	t = 0x00000001
	d = 0x00000010
	l = 0x00000100
	s = 0x00001000
	m = 0x00010000
)

var (
	options = map[string]int64{
		"%T": t,
		"%D": d,
		"%L": l,
		"%S": s,
		"%M": m,
	}
	adapter = strings.NewReplacer(
		"%T", "{{if .TIME}}{{.TIME.Format \"15:04:05\"}}{{end}}",
		"%D", "{{if .TIME}}{{.TIME.Format \"2006-01-02\"}}{{end}}",
		"%L", "{{.LEVEL}}",
		"%S", "{{.SOURCE}}",
		"%M", "{{.MESSAGE}}",
	)
)

// 配置文件格式
type xmlProperty struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type xmlFilter struct {
	Enabled  bool          `xml:"enabled,attr"`
	Type     string        `xml:"type"`
	Level    string        `xml:"level"`
	Property []xmlProperty `xml:"property"`
}

type xmlLoggerConfig struct {
	Filter  []xmlFilter `xml:"filter"`
	Unmerge bool        `xml:"unmerge,attr"`
}
