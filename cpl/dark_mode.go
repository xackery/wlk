//go:build windows
// +build windows

package cpl

import (
	"syscall"
	"unsafe"

	"github.com/xackery/wlk/common"
	"github.com/xackery/wlk/win"
)

// SetDarkModeAllowed is used to allow dark mode. This should be called prior to initializing walk
func SetDarkModeAllowed(value bool) {
	common.SetDarkModeAllowed(value)
}

// IsDarkMode returns true if the user is in dark mode
func IsDarkMode() bool {
	if common.IsDarkModeChecked() {
		return common.IsDarkMode()
	}

	common.SetDarkModeChecked(true)
	var hKey win.HKEY

	ptr, err := syscall.UTF16PtrFromString(`Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`)
	if err != nil {
		return false
	}

	if win.RegOpenKeyEx(
		win.HKEY_CURRENT_USER,
		ptr,
		0,
		win.KEY_READ,
		&hKey) != win.ERROR_SUCCESS {
		return false
	}

	defer win.RegCloseKey(hKey)

	bufSize := uint32(4)
	var val uint32

	valueName, err := syscall.UTF16PtrFromString("AppsUseLightTheme")
	if err != nil {
		return false
	}

	if win.ERROR_SUCCESS != win.RegQueryValueEx(
		hKey,
		valueName,
		nil,
		nil,
		(*byte)(unsafe.Pointer(&val)),
		&bufSize) {
		return false
	}

	common.SetDarkMode(val == 0)
	return common.IsDarkMode()
}
