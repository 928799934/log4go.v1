package log4go

import (
	"strings"
)

// template 模版适配器
const (
	T = 0x00000001
	D = 0x00000010
	L = 0x00000100
	S = 0x00001000
	M = 0x00010000
)

var (
	options = map[string]int64{
		"%T": T,
		"%D": D,
		"%L": L,
		"%S": S,
		"%M": M,
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
