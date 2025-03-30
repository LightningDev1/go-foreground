//go:build darwin

package foreground

import (
	"errors"
	"os/exec"
	"strings"
)

// GetForegroundPID returns the PID of the foreground window.
func GetForegroundPID() (uint32, error) {
	return 0, errors.New("GetForegroundPID() not implemented on darwin")
}

// GetForegroundTitle returns the title of the foreground window.
func GetForegroundTitle() (string, error) {
	cmd := exec.Command("osascript", "-e",
		`tell application "System Events" to tell (first process whose frontmost is true) to return name of window 1`)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
