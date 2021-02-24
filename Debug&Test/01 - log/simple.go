package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./log.txt") // file will be overwriten
	defer logfile.Close()

	// parameters here is io.Writer = our log.txt file
	// prefix with space at end,
	// byte-mask of flags for date format and filename with row number
	// LstdFlags = Ldate + Ltime; Lmicrosends - rise precision
	// Lshortfile - only name w/row; Llongfile - whole path
	logger := log.New(logfile, "crab ", log.LstdFlags|log.Lshortfile)

	logger.Println("Usual")
	logger.Fatalln("Fatal") // prints fatal error and do os.Exit(1)
	logger.Println("End")   // never be executed
}
