//go:build windows

package foreground

import (
	"github.com/BurntSushi/xgb"
	"github.com/robotn/xgb/xproto"
)

// getActiveWindowID returns the X11 window ID of the active window.
func getActiveWindowID() (xproto.Window, error) {
	conn, err := xgb.NewConn()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	setup := xproto.Setup(conn)
	screen := setup.DefaultScreen(conn)
	root := screen.Root

	activeWindow, err := xproto.GetProperty(
		conn,
		false,
		root,
		xproto.Atom("_NET_ACTIVE_WINDOW"),
		xproto.Atom("WINDOW"),
		0,
		(1<<32)-1,
	).Reply()
	if err != nil {
		return 0, err
	}

	return xproto.Window(activeWindow.Value[0]), nil
}

// GetForegroundPID returns the PID of the foreground window.
func GetForegroundPID() (uint32, error) {
	windowID, err := getActiveWindowID()
	if err != nil {
		return 0, err
	}

	conn, err := xgb.NewConn()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	propReply, err := xproto.GetProperty(
		conn,
		false,
		windowID,
		xproto.Atom("_NET_WM_PID"),
		xproto.Atom("CARDINAL"),
		0,
		(1<<32)-1,
	).Reply()
	if err != nil {
		return 0, err
	}

	return propReply.Value32(), nil
}

// GetForegroundTitle returns the title of the foreground window.
func GetForegroundTitle() (string, error) {
	windowID, err := getActiveWindowID()
	if err != nil {
		return "", err
	}

	conn, err := xgb.NewConn()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	propReply, err := xproto.GetProperty(
		conn,
		false,
		windowID,
		xproto.Atom("_NET_WM_NAME"),
		xproto.Atom("UTF8_STRING"),
		0,
		(1<<32)-1,
	).Reply()
	if err != nil {
		return "", err
	}

	return string(propReply.Value), nil
}
