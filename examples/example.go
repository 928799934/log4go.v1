package main

import (
	log "github.com/928799934/log4go.v1"
)

func main() {
	defer log.Close()
	for x := 0; x < 2; x++ {
		if err := log.LoadConfiguration("logformat.xml"); err != nil {
			return
		}
		ll := log.Logger(log.DEBUG)
		log.Debug("debug")
		ll.Println("xxx")
		//log.Trace("trace")

		//ll.Println("xxx")
		//log.Warn("warn")
		//ll.Println("xxx")
		//log.Error("error")
		//log.Error("error")
		//log.Error("error")

		//log.Info("info asdadadlkjadadlkjalkjdalkjdlkjadlkjalkjdalkjdlkjadlkjada")
	}
}
