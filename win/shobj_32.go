// Copyright 2012 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (windows && 386) || (windows && arm)
// +build windows,386 windows,arm

package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func (obj *ITaskbarList3) SetProgressValue(hwnd windows.HWND, current uint32, length uint32) HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.SetProgressValue, 6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		uintptr(current),
		0,
		uintptr(length),
		0)

	return HRESULT(ret)
}
