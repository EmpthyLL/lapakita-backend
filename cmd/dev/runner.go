package main

import (
	"fmt"
	"os"
	"os/exec"
)

func RunDev() error {
	config := AirConfig()

	fmt.Printf("Starting Lapakita backend...\n")
	fmt.Printf("Using Air config: %s\n\n", config)

	if _, err := os.Stat(config); err != nil {
		return fmt.Errorf("Air config not found: %s", config)
	}

	air, err := exec.LookPath("air")
	if err != nil {
		return fmt.Errorf(
			"Air is not installed or not available in PATH; install with: go install github.com/air-verse/air@latest",
		)
	}

	cmd := exec.Command(air, "-c", config)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}