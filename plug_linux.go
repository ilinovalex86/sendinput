package sendinput

func (ke *KeyboardEvent) ShiftPress()   {}
func (ke *KeyboardEvent) ShiftRelease() {}
func (ke *KeyboardEvent) CtrlPress()    {}
func (ke *KeyboardEvent) CtrlRelease()  {}
func (ke *KeyboardEvent) Launching() error {
	return nil
}

func (m *MouseEvent) Move() error {
	return nil
}

func (m *MouseEvent) LClick() error {
	return nil
}

func (m *MouseEvent) RClick() error {
	return nil
}

func (m *MouseEvent) DoubleClick() error {
	return nil
}

func (m *MouseEvent) Drop() error {
	return nil
}

func (m *MouseEvent) WheelUp() error {
	return nil
}

func (m *MouseEvent) WheelDown() error {
	return nil
}
