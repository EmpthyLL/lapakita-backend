package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func RunDev() error {
	config := ".air.toml"

	if runtime.GOOS == "windows" {
		config = ".air.windows.toml"
	}

	fmt.Printf("Starting Lapakita backend...\n")
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Using Air config: %s\n\n", config)

	if _, err := os.Stat(config); err != nil {
		return fmt.Errorf("Air config not found: %s", config)
	}

	air, err := exec.LookPath("air")
	if err != nil {
		return fmt.Errorf(
			"Air is not installed or not available in PATH: %w",
			err,
		)
	}

	cmd := exec.Command(air, "-c", config)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
