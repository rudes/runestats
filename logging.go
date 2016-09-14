package main

import (
	"log"
	"os"
)

// TODO: Transactional SQL Logging

func logIt(t ...interface{}) {
	logFile := "./runestats.log"
	f, _ := os.OpenFile(logFile,
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	log.SetOutput(f)
	log.Println(t...)
}
