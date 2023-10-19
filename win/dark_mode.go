//go:build windows
// +build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/xackery/wlk/common"
)

// IsDarkMode returns true if the user is in dark mode
func IsDarkMode() bool {
	if common.IsDarkModeChecked() {
		return common.IsDarkMode()
	}

	common.SetDarkModeChecked(true)
	var hKey HKEY

	ptr, err := syscall.UTF16PtrFromString(`Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`)
	if err != nil {
		return false
	}

	if RegOpenKeyEx(
		HKEY_CURRENT_USER,
		ptr,
		0,
		KEY_READ,
		&hKey) != ERROR_SUCCESS {
		return false
	}

	defer RegCloseKey(hKey)

	bufSize := uint32(4)
	var val uint32

	valueName, err := syscall.UTF16PtrFromString("AppsUseLightTheme")
	if err != nil {
		return false
	}

	if ERROR_SUCCESS != RegQueryValueEx(
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
