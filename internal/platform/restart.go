package platform

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Restart launches a fresh copy of the current executable and exits.
func Restart() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		return err
	}

	cmd := exec.Command(exe, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Start(); err != nil {
		return err
	}

	os.Exit(0)
	return nil
}
