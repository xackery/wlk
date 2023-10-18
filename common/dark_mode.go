package common

const (
	DarkFormBG            = Color(0x202020)
	DarkFormLighterBG     = Color(0x272727)
	DarkButtonBG          = Color(0x343434)
	DarkTextTitleFG       = Color(0xEEEEEE)
	DarkTextFG            = Color(0xDBDBD1)
	DarkTextLinkFG        = Color(0x99EBFF)
	DarkSelectHoverBG     = Color(0x162430)
	DarkSelectHighlightBG = Color(0x143047)
)

var (
	isDarkMode        bool
	isDarkModeAllowed bool
	isDarkModeChecked bool
)

// IsDarkMode returns true if dark mode is enabled and allowed
func IsDarkMode() bool {
	if !isDarkModeAllowed {
		return false
	}
	return isDarkMode
}

// SetDarkModeAllowed is used to allow dark mode. This should be called prior to initializing walk
func SetDarkModeAllowed(value bool) {
	isDarkModeAllowed = value
}

// SetDarkMode is used to set if dark mode is enabled. This is normally set by the OS automatically, but you can override it too
func SetDarkMode(value bool) {
	isDarkMode = value
}

// IsDarkModeChecked returns true if dark mode has been checked already. This can be ignored as it's set automatically
func IsDarkModeChecked() bool {
	return isDarkModeChecked
}

// SetDarkModeChecked is used to set if dark mode has been checked. This can be ignored as it's set automatically
func SetDarkModeChecked(value bool) {
	isDarkModeChecked = value
}
