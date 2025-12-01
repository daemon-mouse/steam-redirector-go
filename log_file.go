package main

import (
	"io"
	"log"
	"os"
)

var logWriter io.Writer
var logFile *os.File

// startLog opens the log file steam-redirector.log and arranges for all future
// calls to functions in the log package to write to it in addition to stderr.
func startLog() {
	// If this is a first run, we want to erase any existing log file and create
	// it if it doesn't exist.
	logFileFlags := os.O_WRONLY | os.O_CREATE
	if _, has := os.LookupEnv(EnvNoRedirect); has {
		// On a second run, we want to append to the log file rather than
		// clearing it.
		logFileFlags |= os.O_APPEND
	}

	lf, err := os.OpenFile("steam-redirector.log", logFileFlags, 0644)
	if err != nil {
		// It is not a fatal error for the log file to be unwritable, since the user
		// would certainly prefer for the game to still launch if possible. A log file
		// error can happen in legitimate scenarios: for example, a user may want to
		// prevent a game from overwriting an INI file by making the game directory
		// read-only. When the log file can't be opened, we write logs only to stderr.
		defer log.Printf("warn: could not open steam-redirector.log: %v\n", err)
		logWriter = os.Stderr
	} else {
		logWriter, logFile = io.MultiWriter(os.Stderr, lf), lf
	}
	log.SetOutput(logWriter)
	log.SetPrefix("steam-redirector: ")
}

// closeLog closes our log file if it was opened. It logs any I/O errors that
// occur.
func closeLog() {
	if logFile != nil {
		if err := logFile.Close(); err != nil {
			log.Printf("warn: could not close steam-redirector.log: %v\n", err)
		}
	}
}
