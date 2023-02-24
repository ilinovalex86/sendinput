package sendinput

import (
	"time"
	"unsafe"
)

//mouseData is struct for winApi mouse_event.
type mouseData struct {
	dx        int32
	dy        int32
	mouseData int32
	flags     uint32
	time      uint32
	extraInfo uintptr
}

//mouseInput is struct for winApi sendInput.
type mouseInput struct {
	inputType uint32
	md        mouseData
}

func (ms *mouseInput) sendInput() (int, error) {
	ret, _, err := sendInputProc.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(ms)),
		unsafe.Sizeof(*ms),
	)
	return int(ret), err
}

//newMouseInput create winApi struct for sendInput
func newMouseInput(md mouseData) *mouseInput {
	var mi mouseInput
	mi.inputType = 0
	mi.md = md
	return &mi
}

func (m *MouseEvent) Move() error {
	ret, _, err := setCursorPosProc.Call(uintptr(m.CordX), uintptr(m.CordY))
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) LClick() error {
	err := m.Move()
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond)
	mi := newMouseInput(mouseData{flags: 0x0002})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(time.Millisecond)
	mi.md.flags = 0x0004
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) RClick() error {
	err := m.Move()
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond)
	mi := newMouseInput(mouseData{flags: 0x0008})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(time.Millisecond)
	mi.md.flags = 0x0010
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) DoubleClick() error {
	err := m.Move()
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond)
	mi := newMouseInput(mouseData{flags: 0x0002})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(time.Millisecond)
	mi.md.flags = 0x0004
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(time.Millisecond)
	mi.md.flags = 0x0002
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(time.Millisecond)
	mi.md.flags = 0x0004
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) Drop() error {
	mi := newMouseInput(mouseData{flags: 0x0002})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(5 * time.Millisecond)
	err = m.Move()
	if err != nil {
		return err
	}
	time.Sleep(25 * time.Millisecond)
	mi.md.flags = 0x0004
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) WheelUp() error {
	mi := newMouseInput(mouseData{flags: 0x0800, mouseData: 120})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) WheelDown() error {
	mi := newMouseInput(mouseData{flags: 0x0800, mouseData: -120})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}
