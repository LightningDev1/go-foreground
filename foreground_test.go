package foreground

import (
	"testing"
)

func TestForeground(t *testing.T) {
	pid, err := GetForegroundPID()
	if err != nil {
		t.Error("GetForegroundPID():", err)
	}

	t.Log("Foreground PID:", pid)

	title, err := GetForegroundTitle()
	if err != nil {
		t.Error("GetForegroundTitle():", err)
	}

	t.Log("Foreground title:", title)
}
