package platform

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const handoffFlag = "--wait-for-pid="

// HandoffPID returns the PID from --wait-for-pid=N, if present in args.
func HandoffPID(args []string) int {
	for _, arg := range args {
		if strings.HasPrefix(arg, handoffFlag) {
			pid, err := strconv.Atoi(strings.TrimPrefix(arg, handoffFlag))
			if err == nil && pid > 0 {
				return pid
			}
		}
	}
	return 0
}

// WaitForExit blocks until the given process is no longer running.
func WaitForExit(pid int) {
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if !processRunning(pid) {
			return
		}
		time.Sleep(50 * time.Millisecond)
	}
}

// LaunchHandoff starts a new instance that waits for this process to exit before showing UI.
func LaunchHandoff() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		return err
	}

	cmd := exec.Command(exe, fmt.Sprintf("%s%d", handoffFlag, os.Getpid()))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Start()
}
