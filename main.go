package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var MO2PathFile = filepath.Join("modorganizer2", "instance_path.txt")

func main() {
	startLog()
	defer closeLog()

	if err := run(); err != nil {
		log.Fatalf("fatal error: %v\n", err)
	}
}

func run() error {
	var exePath string
	var err error
	if _, has := os.LookupEnv("NO_REDIRECT"); !has {
		envErr := os.Setenv("NO_REDIRECT", "1")
		if envErr != nil {
			log.Printf("warn: failed to set environment variable NO_REDIRECT=1: %v\n", err)
		}
		log.Printf("Reading from file...\n")
		exePath, err = readPathFromFile(MO2PathFile)
	} else {
		log.Printf("Reading original launcher...\n")
		exePath, err = getOriginalLauncher(os.Args[0])
	}
	if err != nil {
		log.Printf("error: could not find executable: %v\n", err)
		log.Printf("Launching explorer...\n")
		cmd := exec.Command("explorer.exe")
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		return cmd.Run()
	}

	fmt.Printf("Launching %s...\n", exePath)
	cmd := exec.Command(exePath, os.Args[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, logWriter, logWriter
	return cmd.Run()
}

func readPathFromFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("could not read %s: %w", filePath, err)
	}
	c := strings.TrimSpace(string(content))
	log.Printf("read '%s' - '%s'\n", filePath, c)
	return c, nil
}

func getOriginalLauncher(redirectorPath string) (string, error) {
	return filepath.Join(filepath.Dir(redirectorPath), "_"+filepath.Base(redirectorPath)), nil
}
