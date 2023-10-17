// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	DWMWA_NCRENDERING_ENABLED            = 1
	DWMWA_NCRENDERING_POLICY             = 2
	DWMWA_TRANSITIONS_FORCEDISABLED      = 3
	DWMWA_ALLOW_NCPAINT                  = 4
	DWMWA_CAPTION_BUTTON_BOUNDS          = 5
	DWMWA_NONCLIENT_RTL_LAYOUT           = 6
	DWMWA_FORCE_ICONIC_REPRESENTATION    = 7
	DWMWA_FLIP3D_POLICY                  = 8
	DWMWA_EXTENDED_FRAME_BOUNDS          = 9
	DWMWA_HAS_ICONIC_BITMAP              = 10
	DWMWA_DISALLOW_PEEK                  = 11
	DWMWA_EXCLUDED_FROM_PEEK             = 12
	DWMWA_CLOAK                          = 13
	DWMWA_CLOAKED                        = 14
	DWMWA_FREEZE_REPRESENTATION          = 15
	DWMWA_PASSIVE_UPDATE_MODE            = 16
	DWMWA_USE_HOSTBACKDROPBRUSH          = 17
	DWMWA_USE_IMMERSIVE_DARK_MODE        = 20
	DWMWA_WINDOW_CORNER_PREFERENCE       = 33
	DWMWA_BORDER_COLOR                   = 34
	DWMWA_CAPTION_COLOR                  = 35
	DWMWA_TEXT_COLOR                     = 36
	DWMWA_VISIBLE_FRAME_BORDER_THICKNESS = 37
)

var (
	dwmapi                    *windows.LazyDLL
	procDwmGetWindowAttribute *windows.LazyProc
	procDwmSetWindowAttribute *windows.LazyProc
)

func init() {
	dwmapi = windows.NewLazySystemDLL("dwmapi.dll")

	procDwmGetWindowAttribute = dwmapi.NewProc("DwmGetWindowAttribute")
	procDwmSetWindowAttribute = dwmapi.NewProc("DwmSetWindowAttribute")
}

func DwmGetWindowAttribute(hwnd HWND, attribute uint32) (value uint32, ret error) {
	size := unsafe.Sizeof(value)
	r0, _, _ := syscall.Syscall6(procDwmGetWindowAttribute.Addr(), 4, uintptr(hwnd), uintptr(attribute), uintptr(value), size, 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func DwmSetWindowAttribute(hwnd HWND, attribute uint32, value uint32) (ret error) {

	size := unsafe.Sizeof(value)
	r0, _, _ := syscall.Syscall6(procDwmSetWindowAttribute.Addr(), 4, uintptr(hwnd), uintptr(attribute), uintptr(value), size, 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}
