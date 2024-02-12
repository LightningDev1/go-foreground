//go:build windows

package foreground

import (
	"syscall"
	"unsafe"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	getForegroundWindow      = user32.NewProc("GetForegroundWindow")
	getWindowThreadProcessId = user32.NewProc("GetWindowThreadProcessId")
	getWindowTextW           = user32.NewProc("GetWindowTextW")
)

// GetForegroundPID returns the PID of the foreground window.
func GetForegroundPID() (uint32, error) {
	foregroundWindow, _, _ := getForegroundWindow.Call()

	var pid uint32
	_, _, err := getWindowThreadProcessId.Call(foregroundWindow, uintptr(unsafe.Pointer(&pid)))
	if err != nil && err.Error() != "The operation completed successfully." {
		return 0, err
	}

	return pid, nil
}

// GetForegroundTitle returns the title of the foreground window.
func GetForegroundTitle() (string, error) {
	const nMaxCount = 1024

	foregroundWindow, _, _ := getForegroundWindow.Call()

	var title [nMaxCount]uint16
	_, _, err := getWindowTextW.Call(foregroundWindow, uintptr(unsafe.Pointer(&title[0])), nMaxCount)
	if err != nil && err.Error() != "The operation completed successfully." {
		return "", err
	}

	return syscall.UTF16ToString(title[:]), nil
}
