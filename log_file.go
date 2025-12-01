package main

import (
	"io"
	"log"
	"os"
)

var logWriter io.Writer
var logFile *os.File

func startLog() {
	logFileFlags := os.O_WRONLY | os.O_CREATE
	if _, has := os.LookupEnv("NO_REDIRECT"); has {
		logFileFlags |= os.O_APPEND
	}

	lf, err := os.OpenFile("steam-redirector.log", logFileFlags, 0644)
	if err != nil {
		defer log.Printf("warn: could not open steam-redirector.log: %v\n", err)
		logWriter = os.Stderr
	} else {
		logWriter, logFile = io.MultiWriter(os.Stderr, lf), lf
	}
	log.SetOutput(logWriter)
	log.SetPrefix("steam-redirector: ")
}

func closeLog() {
	if logFile != nil {
		if err := logFile.Close(); err != nil {
			log.Printf("warn: could not close steam-redirector.log: %v\n", err)
		}
	}
}
