//go:build windows
// +build windows

package win

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// DwmGetWindowAttribute gets the value of a window attribute
func DwmGetWindowAttribute(hwnd windows.HWND, attribute uint32) (value uint32, ret error) {
	err := windows.DwmGetWindowAttribute(hwnd, attribute, unsafe.Pointer(&value), 4)
	return value, err
}

// DwmSetWindowAttribute sets the value of a window attribute
func DwmSetWindowAttribute(hwnd windows.HWND, attribute uint32, value uint32) (ret error) {
	return windows.DwmSetWindowAttribute(hwnd, attribute, unsafe.Pointer(&value), 4)
}
