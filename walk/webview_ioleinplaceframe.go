// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package walk

import (
	"syscall"

	"github.com/xackery/wlk/win"
)

var webViewIOleInPlaceFrameVtbl *win.IOleInPlaceFrameVtbl

func init() {
	AppendToWalkInit(func() {
		webViewIOleInPlaceFrameVtbl = &win.IOleInPlaceFrameVtbl{
			QueryInterface:       syscall.NewCallback(webView_IOleInPlaceFrame_QueryInterface),
			AddRef:               syscall.NewCallback(webView_IOleInPlaceFrame_AddRef),
			Release:              syscall.NewCallback(webView_IOleInPlaceFrame_Release),
			GetWindow:            syscall.NewCallback(webView_IOleInPlaceFrame_GetWindow),
			ContextSensitiveHelp: syscall.NewCallback(webView_IOleInPlaceFrame_ContextSensitiveHelp),
			GetBorder:            syscall.NewCallback(webView_IOleInPlaceFrame_GetBorder),
			RequestBorderSpace:   syscall.NewCallback(webView_IOleInPlaceFrame_RequestBorderSpace),
			SetBorderSpace:       syscall.NewCallback(webView_IOleInPlaceFrame_SetBorderSpace),
			SetActiveObject:      syscall.NewCallback(webView_IOleInPlaceFrame_SetActiveObject),
			InsertMenus:          syscall.NewCallback(webView_IOleInPlaceFrame_InsertMenus),
			SetMenu:              syscall.NewCallback(webView_IOleInPlaceFrame_SetMenu),
			RemoveMenus:          syscall.NewCallback(webView_IOleInPlaceFrame_RemoveMenus),
			SetStatusText:        syscall.NewCallback(webView_IOleInPlaceFrame_SetStatusText),
			EnableModeless:       syscall.NewCallback(webView_IOleInPlaceFrame_EnableModeless),
			TranslateAccelerator: syscall.NewCallback(webView_IOleInPlaceFrame_TranslateAccelerator),
		}
	})
}

type webViewIOleInPlaceFrame struct {
	win.IOleInPlaceFrame
	webView *WebView
}

func webView_IOleInPlaceFrame_QueryInterface(inPlaceFrame *webViewIOleInPlaceFrame, riid win.REFIID, ppvObj *uintptr) uintptr {
	return win.E_NOTIMPL
}

func webView_IOleInPlaceFrame_AddRef(inPlaceFrame *webViewIOleInPlaceFrame) uintptr {
	return 1
}

func webView_IOleInPlaceFrame_Release(inPlaceFrame *webViewIOleInPlaceFrame) uintptr {
	return 1
}

func webView_IOleInPlaceFrame_GetWindow(inPlaceFrame *webViewIOleInPlaceFrame, lphwnd *win.HWND) uintptr {
	*lphwnd = inPlaceFrame.webView.hWnd

	return win.S_OK
}

func webView_IOleInPlaceFrame_ContextSensitiveHelp(inPlaceFrame *webViewIOleInPlaceFrame, fEnterMode win.BOOL) uintptr {
	return win.E_NOTIMPL
}

func webView_IOleInPlaceFrame_GetBorder(inPlaceFrame *webViewIOleInPlaceFrame, lprectBorder *win.RECT) uintptr {
	return win.E_NOTIMPL
}

func webView_IOleInPlaceFrame_RequestBorderSpace(inPlaceFrame *webViewIOleInPlaceFrame, pborderwidths uintptr) uintptr {
	return win.E_NOTIMPL
}

func webView_IOleInPlaceFrame_SetBorderSpace(inPlaceFrame *webViewIOleInPlaceFrame, pborderwidths uintptr) uintptr {
	return win.E_NOTIMPL
}

func webView_IOleInPlaceFrame_SetActiveObject(inPlaceFrame *webViewIOleInPlaceFrame, pActiveObject uintptr, pszObjName *uint16) uintptr {
	return win.S_OK
}

func webView_IOleInPlaceFrame_InsertMenus(inPlaceFrame *webViewIOleInPlaceFrame, hmenuShared win.HMENU, lpMenuWidths uintptr) uintptr {
	return win.E_NOTIMPL
}

func webView_IOleInPlaceFrame_SetMenu(inPlaceFrame *webViewIOleInPlaceFrame, hmenuShared win.HMENU, holemenu win.HMENU, hwndActiveObject win.HWND) uintptr {
	return win.S_OK
}

func webView_IOleInPlaceFrame_RemoveMenus(inPlaceFrame *webViewIOleInPlaceFrame, hmenuShared win.HMENU) uintptr {
	return win.E_NOTIMPL
}

func webView_IOleInPlaceFrame_SetStatusText(inPlaceFrame *webViewIOleInPlaceFrame, pszStatusText *uint16) uintptr {
	return win.S_OK
}

func webView_IOleInPlaceFrame_EnableModeless(inPlaceFrame *webViewIOleInPlaceFrame, fEnable win.BOOL) uintptr {
	return win.S_OK
}

func webView_IOleInPlaceFrame_TranslateAccelerator(inPlaceFrame *webViewIOleInPlaceFrame, lpmsg *win.MSG, wID uint32) uintptr {
	return win.E_NOTIMPL
}
