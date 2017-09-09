## Log4go
This package is a replacement logging package which will be both a drop-in replacement for and a significant extension of the built-in logging functionality in Go.
## Option:
```sh
<!--  unmerge (true|false) - 是否合并数据(true:数据合并将error显示到info等文件中) -->
<logging unmerge="true">

    <!-- enabled (true|false) - 是否启用本过滤器 -->
    <filter enabled="true">

        <!-- 输出类型 (console|file) -->
        <type>file</type>

        <!-- 输出等级 (DEBUG|TRACE|INFO|WARNING|ERROR) -->
        <level>INFO</level>

        <!--
             输出格式:
             %T - 时间 (15:04:05)
             %D - 日期 (2006/01/02)
             %L - 等级 (DEBG, TRAC, WARN, EROR)
             %S - 源文件路径
             %M - 信息
             例: "[%D %T] [%L] (%S) %M"
        -->
        <property name="format">[%D %T] [%L] [%S] %M</property>

        <!-- 文件存储路径 -->
        <property name="filename">E:\code_source\page_game\go\trunk\src\common\log4go.v1\examples\info.log</property>

        <!-- 文件切割体积 (k/m/g|K/M/G) -->
        <property name="maxsize">0M</property> <!-- \d+[KMG]? Suffixes are in terms of 2**10 -->

        <!-- 延迟写入 (s/m/h|S/M/H) -->
        <property name="delay">15s</property> <!-- \d+[SMH]?  Suffixes are in terms of 60 -->

    </filter>

    <!--  enabled (true|false) - 是否启用本过滤器 -->
    <filter enabled="true">

        <!-- 输出类型 (console|file) -->
        <type>console</type>

        <!-- 输出等级 (DEBUG|TRACE|INFO|WARNING|ERROR) -->
        <level>DEBUG</level>

        <!--
             输出格式:
             %T - 时间 (15:04:05)
             %D - 日期 (2006/01/02)
             %L - 等级 (DEBG, TRAC, WARN, EROR)
             %S - 源文件路径
             %M - 信息
             例: "[%D %T] [%L] (%S) %M"
        -->
        <property name="format">[%D] [%L] [%S] %M</property>

    </filter>
</logging>
```
## Usage:
---------
```go
package main

import (
	log "git.liebaopay.com/pigs/public/log4go"
)

func main() {
	defer log.Close()
	for x := 0; x < 2; x++ {
		if err := log.LoadConfiguration("logformat.xml"); err != nil {
			return
		}
		ll := log.Logger(log.ERROR)
		log.Debug("debug")
		ll.Println("xxx")
		log.Trace("trace")
		log.Warn("warn")
		log.Error("error")
		log.Info("info asdadadlkjadadlkjalkjdalkjdlkjadlkjalkjdalkjdlkjadlkjada")
	}
}
```
