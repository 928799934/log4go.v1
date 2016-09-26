package main

import (
	log "log4go.v1"
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
