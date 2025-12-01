package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// main arranges for our log file to open, handles fatal errors, and exits with
// the appropriate code.
func main() {
	exitCode := 0

	startLog()
	defer func() {
		closeLog()
		os.Exit(exitCode)
	}()

	if err := run(); err != nil {
		log.Printf("fatal error: %v\n", err)
		exitCode = 1
	}
}

// run is the primary logic of the program. It determines which executable to run,
// and then launches it with the appropriate arguments.
func run() error {
	var exePath string
	var err error

	// This program is run twice on a usual game launch: once by the game
	// launcher (Steam, Heroic, etc.), and then by ModOrganizer 2. We set an
	// environment variable on the first run so that we can tell which one we're
	// on.
	if _, has := os.LookupEnv(EnvNoRedirect); !has {
		envErr := os.Setenv(EnvNoRedirect, "1")
		if envErr != nil {
			log.Printf("warn: failed to set environment variable %s=1: %v\n", EnvNoRedirect, err)
		}

		// This is our first launch. Launch ModOrganizer 2.
		log.Printf("Reading from file...\n")
		exePath, err = readPathFromFile(MO2PathFile)
	} else {
		// This is our second launch. Launch the actual game.
		log.Printf("Reading original launcher...\n")
		exePath, err = getOriginalLauncher(os.Args[0])
	}
	if err != nil {
		// Launch explorer as a fallback in case of error.
		log.Printf("error: could not find executable: %v\n", err)
		log.Printf("Launching explorer...\n")
		cmd := exec.Command("explorer.exe")
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		return cmd.Run()
	}

	// Scan for the first argument that appears to be intended for ModOrganizer
	// 2, and use that as the sole argument to the subprocess.
	var args []string
scanArgs:
	for _, arg := range os.Args[1:] {
		switch {
		case strings.HasPrefix(arg, SchemeNXM),
			strings.HasPrefix(arg, SchemeMO2Shortcut),
			arg == MO2ArgPick:
			args = []string{arg}
			break scanArgs
		}
	}
	fmt.Printf("Launching %s...\n", exePath)
	cmd := exec.Command(exePath, args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, logWriter, logWriter
	return cmd.Run()
}

// readPathFromFile opens the given file and returns its contents.
func readPathFromFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("could not read %s: %w", filePath, err)
	}
	c := strings.TrimSpace(string(content))
	log.Printf("read '%s' - '%s'\n", filePath, c)
	return c, nil
}

// getOriginalLauncher infers the path to the original launcher executable.
func getOriginalLauncher(redirectorPath string) (string, error) {
	return filepath.Join(filepath.Dir(redirectorPath), "_"+filepath.Base(redirectorPath)), nil
}
