package sendinput

import (
	"errors"
	"fmt"
	"time"
	"unsafe"
)

const (
	shift = 0x10
	ctrl  = 0x11
	alt   = 0x12
	win   = 0x5B
	keyUp = 0x0002
)

//keyboardData is struct for winApi keybd_event.
type keyboardData struct {
	vk        uint16
	scan      uint16
	flags     uint32
	time      uint32
	extraInfo uintptr
}

//keyboardInput is struct for winApi sendInput.
type keyboardInput struct {
	inputType uint32
	kd        keyboardData
	padding   uint64
}

//newKeyboardInput create winApi struct for sendInput
func newKeyboardInput(kd keyboardData) *keyboardInput {
	var ki keyboardInput
	ki.inputType = 1 //INPUT_KEYBOARD
	ki.kd = kd
	return &ki
}

func (ki *keyboardInput) sendInput() (int, error) {
	ret, _, err := sendInputProc.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(ki)),
		unsafe.Sizeof(*ki),
	)
	return int(ret), err
}

func downKey(key uint16) error {
	ki := newKeyboardInput(keyboardData{vk: key})
	ret, err := ki.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func upKey(key uint16) error {
	ki := newKeyboardInput(keyboardData{vk: key, flags: keyUp})
	ret, err := ki.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func (ke *KeyboardEvent) ShiftPress()   { _ = downKey(shift) }
func (ke *KeyboardEvent) ShiftRelease() { _ = upKey(shift) }
func (ke *KeyboardEvent) CtrlPress()    { _ = downKey(ctrl) }
func (ke *KeyboardEvent) CtrlRelease()  { _ = upKey(ctrl) }

//Launching use for KeyboardEvent struct
func (ke *KeyboardEvent) Launching() error {
	if key, ok := javaScriptToUint16[ke.JavaScriptCode]; ok {
		if ke.Ctrl {
			downKey(ctrl)
			defer upKey(ctrl)
		}
		if ke.Alt {
			downKey(alt)
			defer upKey(alt)
		}
		if ke.Win {
			downKey(win)
			defer upKey(win)
		}
		if ke.Shift {
			downKey(shift)
			defer upKey(shift)
		}
		err := downKey(key)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond)
		err = upKey(key)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 5)
		return nil
	}
	return errors.New("not in map")
}

func DownKey(javaScriptCode string) error {
	if key, ok := javaScriptToUint16[javaScriptCode]; ok {
		err := downKey(key)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond)
		return nil
	}
	return errors.New("not in map")
}

func UpKey(javaScriptCode string) error {
	if key, ok := javaScriptToUint16[javaScriptCode]; ok {
		err := upKey(key)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 5)
		return nil
	}
	return errors.New("not in map")
}

func StringInput(str string) error {
	for _, r := range str {
		symbol := string(r)
		if javaScriptCode, ok := engStrToJavaScriptCode[symbol]; ok {
			if languageSelect != "eng" {
				LanguageChange.Launching()
				languageSelect = "eng"
				time.Sleep(time.Second * 1)
			}
			ke := KeyboardEvent{JavaScriptCode: javaScriptCode}
			err := ke.Launching()
			if err != nil {
				return err
			}
			continue
		}
		if javaScriptCode, ok := engStrShiftToJavaScriptCode[symbol]; ok {
			if languageSelect != "eng" {
				LanguageChange.Launching()
				languageSelect = "eng"
				time.Sleep(time.Second * 1)
			}
			ke := KeyboardEvent{Shift: true, JavaScriptCode: javaScriptCode}
			err := ke.Launching()
			if err != nil {
				return err
			}
			continue
		}
		if javaScriptCode, ok := rusStrToJavaScriptCode[symbol]; ok {
			if languageSelect == "eng" {
				LanguageChange.Launching()
				languageSelect = "rus"
				time.Sleep(time.Second * 1)
			}
			ke := KeyboardEvent{JavaScriptCode: javaScriptCode}
			err := ke.Launching()
			if err != nil {
				return err
			}
			continue
		}
		if javaScriptCode, ok := rusStrShiftToJavaScriptCode[symbol]; ok {
			if languageSelect == "eng" {
				LanguageChange.Launching()
				languageSelect = "rus"
				time.Sleep(time.Second * 1)
			}
			ke := KeyboardEvent{Shift: true, JavaScriptCode: javaScriptCode}
			err := ke.Launching()
			if err != nil {
				return err
			}
			continue
		}
		fmt.Println("not in map: ", symbol, " code: ", r)
	}
	return nil
}

var javaScriptToUint16 = map[string]uint16{

	"Escape": 0x1B,
	"F1":     0x70,
	"F2":     0x71,
	"F3":     0x72,
	"F4":     0x73,
	"F5":     0x74,
	"F6":     0x75,
	"F7":     0x76,
	"F8":     0x77,
	"F9":     0x78,
	"F10":    0x79,
	"F11":    0x7a,
	"F12":    0x7b,

	"Backquote": 0xC0,
	"Digit1":    0x31,
	"Digit2":    0x32,
	"Digit3":    0x33,
	"Digit4":    0x34,
	"Digit5":    0x35,
	"Digit6":    0x36,
	"Digit7":    0x37,
	"Digit8":    0x38,
	"Digit9":    0x39,
	"Digit0":    0x30,
	"Minus":     0xBD,
	"Equal":     0xBB,
	"Backspace": 0x08,

	"Tab":          0x09,
	"KeyQ":         0x51,
	"KeyW":         0x57,
	"KeyE":         0x45,
	"KeyR":         0x52,
	"KeyT":         0x54,
	"KeyY":         0x59,
	"KeyU":         0x55,
	"KeyI":         0x49,
	"KeyO":         0x4F,
	"KeyP":         0x50,
	"BracketLeft":  0xDB,
	"BracketRight": 0xDD,
	"Backslash":    0xDC,

	"CapsLock":  0x14,
	"KeyA":      0x41,
	"KeyS":      0x53,
	"KeyD":      0x44,
	"KeyF":      0x46,
	"KeyG":      0x47,
	"KeyH":      0x48,
	"KeyJ":      0x4A,
	"KeyK":      0x4B,
	"KeyL":      0x4C,
	"Semicolon": 0xBA,
	"Quote":     0xDE,
	"Enter":     0x0D,

	"ShiftLeft":  0xA0,
	"KeyZ":       0x5A,
	"KeyX":       0x58,
	"KeyC":       0x43,
	"KeyV":       0x56,
	"KeyB":       0x42,
	"KeyN":       0x4E,
	"KeyM":       0x4D,
	"Comma":      0xBC,
	"Period":     0xBE,
	"Slash":      0xBF,
	"ShiftRight": 0xA1,

	"ControlLeft":  0xA2,
	"MetaLeft":     0x5B,
	"AltLeft":      0xA4,
	"Space":        0x20,
	"AltRight":     0xA5,
	"MetaRight":    0x5C,
	"ContextMenu":  0x02,
	"ControlRight": 0xA3,

	"Insert":   0x2D,
	"Delete":   0x2E,
	"Home":     0x24,
	"End":      0x23,
	"PageUp":   0x21,
	"PageDown": 0x22,

	"Pause":       0x13,
	"PrintScreen": 0x2c,
	"NumLock":     0x90,
	"ScrollLock":  0x91,

	"ArrowUp":    0x26,
	"ArrowDown":  0x28,
	"ArrowRight": 0x27,
	"ArrowLeft":  0x25,

	"Numpad1":        0x31,
	"Numpad2":        0x32,
	"Numpad3":        0x33,
	"Numpad4":        0x34,
	"Numpad5":        0x35,
	"Numpad6":        0x36,
	"Numpad7":        0x37,
	"Numpad8":        0x38,
	"Numpad9":        0x39,
	"Numpad0":        0x30,
	"NumpadDecimal":  0x6E,
	"NumpadEnter":    0x0D,
	"NumpadAdd":      0x6B,
	"NumpadSubtract": 0x6D,
	"NumpadMultiply": 0x6A,
	"NumpadDivide":   0x6F,
}

var engStrToJavaScriptCode = map[string]string{
	" ": "Space",
	"`": "Backquote",
	"1": "Digit1",
	"2": "Digit2",
	"3": "Digit3",
	"4": "Digit4",
	"5": "Digit5",
	"6": "Digit6",
	"7": "Digit7",
	"8": "Digit8",
	"9": "Digit9",
	"0": "Digit0",
	"-": "Minus",
	"=": "Equal",

	"q":  "KeyQ",
	"w":  "KeyW",
	"e":  "KeyE",
	"r":  "KeyR",
	"t":  "KeyT",
	"y":  "KeyY",
	"u":  "KeyU",
	"i":  "KeyI",
	"o":  "KeyO",
	"p":  "KeyP",
	"[":  "BracketLeft",
	"]":  "BracketRight",
	"\\": "Backslash",

	"a": "KeyA",
	"s": "KeyS",
	"d": "KeyD",
	"f": "KeyF",
	"g": "KeyG",
	"h": "KeyH",
	"j": "KeyJ",
	"k": "KeyK",
	"l": "KeyL",
	";": "Semicolon",
	"'": "Quote",

	"z": "KeyZ",
	"x": "KeyX",
	"c": "KeyC",
	"v": "KeyV",
	"b": "KeyB",
	"n": "KeyN",
	"m": "KeyM",
	",": "Comma",
	".": "Period",
	"/": "Slash",
}

var engStrShiftToJavaScriptCode = map[string]string{
	"~": "Backquote",
	"!": "Digit1",
	"@": "Digit2",
	"#": "Digit3",
	"$": "Digit4",
	"%": "Digit5",
	"^": "Digit6",
	"&": "Digit7",
	"*": "Digit8",
	"(": "Digit9",
	")": "Digit0",
	"_": "Minus",
	"+": "Equal",

	"Q": "KeyQ",
	"W": "KeyW",
	"E": "KeyE",
	"R": "KeyR",
	"T": "KeyT",
	"Y": "KeyY",
	"U": "KeyU",
	"I": "KeyI",
	"O": "KeyO",
	"P": "KeyP",
	"{": "BracketLeft",
	"}": "BracketRight",
	"|": "Backslash",

	"A":  "KeyA",
	"S":  "KeyS",
	"D":  "KeyD",
	"F":  "KeyF",
	"G":  "KeyG",
	"H":  "KeyH",
	"J":  "KeyJ",
	"K":  "KeyK",
	"L":  "KeyL",
	":":  "Semicolon",
	"\"": "Quote",

	"Z": "KeyZ",
	"X": "KeyX",
	"C": "KeyC",
	"V": "KeyV",
	"B": "KeyB",
	"N": "KeyN",
	"M": "KeyM",
	"<": "Comma",
	">": "Period",
	"?": "Slash",
}

var rusStrToJavaScriptCode = map[string]string{
	"ё": "Backquote",

	"й": "KeyQ",
	"ц": "KeyW",
	"у": "KeyE",
	"к": "KeyR",
	"е": "KeyT",
	"н": "KeyY",
	"г": "KeyU",
	"ш": "KeyI",
	"щ": "KeyO",
	"з": "KeyP",
	"х": "BracketLeft",
	"ъ": "BracketRight",

	"ф": "KeyA",
	"ы": "KeyS",
	"в": "KeyD",
	"а": "KeyF",
	"п": "KeyG",
	"р": "KeyH",
	"о": "KeyJ",
	"л": "KeyK",
	"д": "KeyL",
	"ж": "Semicolon",
	"э": "Quote",

	"я": "KeyZ",
	"ч": "KeyX",
	"с": "KeyC",
	"м": "KeyV",
	"и": "KeyB",
	"т": "KeyN",
	"ь": "KeyM",
	"б": "Comma",
	"ю": "Period",
}

var rusStrShiftToJavaScriptCode = map[string]string{
	"Ё": "Backquote",
	"№": "Digit3",

	"Й": "KeyQ",
	"Ц": "KeyW",
	"У": "KeyE",
	"К": "KeyR",
	"Е": "KeyT",
	"Н": "KeyY",
	"Г": "KeyU",
	"Ш": "KeyI",
	"Щ": "KeyO",
	"З": "KeyP",
	"Х": "BracketLeft",
	"Ъ": "BracketRight",

	"Ф": "KeyA",
	"Ы": "KeyS",
	"В": "KeyD",
	"А": "KeyF",
	"П": "KeyG",
	"Р": "KeyH",
	"О": "KeyJ",
	"Л": "KeyK",
	"Д": "KeyL",
	"Ж": "Semicolon",
	"Э": "Quote",

	"Я": "KeyZ",
	"Ч": "KeyX",
	"С": "KeyC",
	"М": "KeyV",
	"И": "KeyB",
	"Т": "KeyN",
	"Ь": "KeyM",
	"Б": "Comma",
	"Ю": "Period",
}
