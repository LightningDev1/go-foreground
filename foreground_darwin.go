//go:build darwin

package foreground

import (
	"errors"

	"github.com/progrium/darwinkit/macos/appkit"
)

// GetForegroundPID returns the PID of the foreground window.
func GetForegroundPID() (uint32, error) {
	return 0, errors.New("GetForegroundPID() not implemented on darwin")
}

// GetForegroundTitle returns the title of the foreground window.
func GetForegroundTitle() (string, error) {
	workspace := appkit.Workspace_SharedWorkspace()
	frontApp := workspace.FrontmostApplication()

	return frontApp.LocalizedName(), nil
}
