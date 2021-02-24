package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	createLogger2()
}

func createLogger1() {
	priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	flags := log.Ldate | log.Lshortfile
	logger, err := syslog.NewLogger(priority, flags)
	if err != nil {
		fmt.Printf("Can't attach to syslog: %s", err)
		return
	}
	logger.Println("Test log message")
}

func createLogger2() {
	logger, err := syslog.New(syslog.LOG_LOCAL3, "elephantus")
	if err != nil {
		panic("Can't attach to syslog")
	}
	defer logger.Close()

	logger.Debug("Debug msg")
	logger.Notice("Notice msg")
	logger.Warning("Warning msg")
	logger.Alert("Alert msg")
}
