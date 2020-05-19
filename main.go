package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatalf("Fatal error reading config file: %s\n", err)
	}

	protoc := GenProtoString(config)

	if err := verifyDirectories(config); err != nil {
		log.Fatalf("Could not create output directories: %v", err)
	}

	if config.Debug {
		log.Printf("Executing protoc command: %v", protoc)
	}

	parts := strings.Split(strings.TrimSpace(protoc), " ")

	cmd := exec.Command(parts[0], parts[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute protoc command: %v", err)
	}

}

func verifyDirectories(config *Config) error {
	if err := createDirIfMissing(config.Output); err != nil {
		return err
	}

	for _, plugin := range config.Plugins {
		if len(plugin.Output) > 0 {
			if err := createDirIfMissing(plugin.Output); err != nil {
				return err
			}
		}
	}
	return nil
}

func createDirIfMissing(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
