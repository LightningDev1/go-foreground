//go:build linux

package foreground

import (
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

// getActiveWindowID returns the X11 window ID of the active window.
func getActiveWindowID() (xproto.Window, error) {
	conn, err := xgb.NewConn()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	setup := xproto.Setup(conn)
	root := setup.DefaultScreen(conn).Root

	aname := "_NET_ACTIVE_WINDOW"
	activeAtom, err := xproto.InternAtom(conn, true, uint16(len(aname)), aname).Reply()
	if err != nil {
		return 0, err
	}

	reply, err := xproto.GetProperty(conn, false, root, activeAtom.Atom,
		xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
	if err != nil {
		return 0, err
	}

	return xproto.Window(xgb.Get32(reply.Value)), nil
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

	aname := "_NET_WM_PID"
	pidAtom, err := xproto.InternAtom(conn, true, uint16(len(aname)), aname).Reply()
	if err != nil {
		return 0, err
	}

	reply, err := xproto.GetProperty(conn, false, windowID, pidAtom.Atom,
		xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
	if err != nil {
		return 0, err
	}

	return xgb.Get32(reply.Value), nil
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

	aname := "_NET_WM_NAME"
	nameAtom, err := xproto.InternAtom(conn, true, uint16(len(aname)), aname).Reply()
	if err != nil {
		return "", err
	}

	reply, err := xproto.GetProperty(conn, false, windowID, nameAtom.Atom,
		xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
	if err != nil {
		return "", err
	}

	return string(reply.Value), nil
}
