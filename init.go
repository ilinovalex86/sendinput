package sendinput

//MouseEvent is type for simple use mouse events.
type MouseEvent struct {
	CordX int
	CordY int
}

var LanguageChange = KeyboardEvent{Win: true, JavaScriptCode: "Space"}
var languageSelect = "eng"

//KeyboardEvent is type for simple use keyboard events. Use method Launching
type KeyboardEvent struct {
	Ctrl           bool
	Shift          bool
	Win            bool
	Alt            bool
	JavaScriptCode string
}
