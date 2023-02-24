package sendinput

import "syscall"

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	sendInputProc    = user32.NewProc("SendInput")
	setCursorPosProc = user32.NewProc("SetCursorPos")
)
