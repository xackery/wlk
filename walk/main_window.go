// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package walk

import (
	"unsafe"

	"github.com/xackery/wlk/common"
	"github.com/xackery/wlk/win"
)

const mainWindowWindowClass = `\o/ Wlk_MainWindow_Class \o/`

func init() {
	AppendToWalkInit(func() {
		MustRegisterWindowClass(mainWindowWindowClass)
	})
}

// MainWindowCfg is a struct for MainWindow configuration
type MainWindowCfg struct {
	Name   string
	Bounds Rectangle
}

// MainWindow represents a top-level window, such as the main window of an application.
type MainWindow struct {
	FormBase
	windowPlacement *win.WINDOWPLACEMENT
	menu            *Menu
	toolBar         *ToolBar
	statusBar       *StatusBar
}

// NewMainWindow creates a new main window with default settings
func NewMainWindow() (*MainWindow, error) {
	return NewMainWindowWithName("")
}

// NewMainWindowWithName creates a new main window with a name
func NewMainWindowWithName(name string) (*MainWindow, error) {
	return NewMainWindowWithCfg(&MainWindowCfg{Name: name})
}

// NewMainWindowWithCfg creates a new main window with a config
func NewMainWindowWithCfg(cfg *MainWindowCfg) (*MainWindow, error) {
	mw := new(MainWindow)
	mw.SetName(cfg.Name)

	if err := initWindowWithCfg(&windowCfg{
		Window:    mw,
		ClassName: mainWindowWindowClass,
		Style:     win.WS_OVERLAPPEDWINDOW,
		ExStyle:   win.WS_EX_CONTROLPARENT,
		Bounds:    cfg.Bounds,
	}); err != nil {
		return nil, err
	}

	succeeded := false
	defer func() {
		if !succeeded {
			mw.Dispose()
		}
	}()

	mw.SetPersistent(true)

	var err error

	if mw.menu, err = newMenuBar(mw); err != nil {
		return nil, err
	}
	if !win.SetMenu(mw.hWnd, mw.menu.hMenu) {
		return nil, lastError("SetMenu")
	}

	tb, err := NewToolBar(mw)
	if err != nil {
		return nil, err
	}
	mw.SetToolBar(tb)

	if mw.statusBar, err = NewStatusBar(mw); err != nil {
		return nil, err
	}
	mw.statusBar.parent = nil
	mw.Children().Remove(mw.statusBar)
	mw.statusBar.parent = mw
	win.SetParent(mw.statusBar.hWnd, mw.hWnd)
	mw.statusBar.visibleChangedPublisher.event.Attach(func() {
		mw.SetBoundsPixels(mw.BoundsPixels())
	})

	if IsDarkMode() {
		bgBrush, err := NewSolidColorBrush(common.DarkFormBG)
		if err != nil {
			return nil, err
		}
		mw.SetBackground(bgBrush)
	}
	succeeded = true

	return mw, nil
}

// Menu returns the menu of the main window.
func (mw *MainWindow) Menu() *Menu {
	return mw.menu
}

// ToolBar returns the tool bar of the main window.
func (mw *MainWindow) ToolBar() *ToolBar {
	return mw.toolBar
}

// SetToolBar sets the tool bar of the main window.
func (mw *MainWindow) SetToolBar(tb *ToolBar) {
	if mw.toolBar != nil {
		win.SetParent(mw.toolBar.hWnd, 0)
	}

	if tb != nil {
		parent := tb.parent
		tb.parent = nil
		parent.Children().Remove(tb)
		tb.parent = mw
		win.SetParent(tb.hWnd, mw.hWnd)
	}

	mw.toolBar = tb
}

// StatusBar returns the status bar of the main window.
func (mw *MainWindow) StatusBar() *StatusBar {
	return mw.statusBar
}

// ClientBoundsPixels returns the bounds of the client area in pixels.
func (mw *MainWindow) ClientBoundsPixels() Rectangle {
	bounds := mw.FormBase.ClientBoundsPixels()

	if mw.toolBar != nil && mw.toolBar.Actions().Len() > 0 {
		tlbBounds := mw.toolBar.BoundsPixels()

		bounds.Y += tlbBounds.Height
		bounds.Height -= tlbBounds.Height
	}

	if mw.statusBar.Visible() {
		bounds.Height -= mw.statusBar.HeightPixels()
	}

	return bounds
}

// SetVisible sets the visibility of the main window
func (mw *MainWindow) SetVisible(visible bool) {
	if visible {
		win.DrawMenuBar(mw.hWnd)

		mw.clientComposite.RequestLayout()
	}

	mw.FormBase.SetVisible(visible)
}

func (mw *MainWindow) applyFont(font *Font) {
	mw.FormBase.applyFont(font)

	if mw.toolBar != nil {
		mw.toolBar.applyFont(font)
	}

	if mw.statusBar != nil {
		mw.statusBar.applyFont(font)
	}
}

// IsFullscreen returns true when the main window is full screen mode, otherwise false
func (mw *MainWindow) IsFullscreen() bool {
	return win.GetWindowLong(mw.hWnd, win.GWL_STYLE)&win.WS_OVERLAPPEDWINDOW == 0
}

// SetFullscreen sets the main window to full screen mode
func (mw *MainWindow) SetFullscreen(fullscreen bool) error {
	if fullscreen == mw.IsFullscreen() {
		return nil
	}

	if fullscreen {
		var mi win.MONITORINFO
		mi.CbSize = uint32(unsafe.Sizeof(mi))

		if mw.windowPlacement == nil {
			mw.windowPlacement = new(win.WINDOWPLACEMENT)
		}

		if !win.GetWindowPlacement(mw.hWnd, mw.windowPlacement) {
			return lastError("GetWindowPlacement")
		}
		if !win.GetMonitorInfo(win.MonitorFromWindow(
			mw.hWnd, win.MONITOR_DEFAULTTOPRIMARY), &mi) {

			return newError("GetMonitorInfo")
		}

		if err := mw.ensureStyleBits(win.WS_OVERLAPPEDWINDOW, false); err != nil {
			return err
		}

		if r := mi.RcMonitor; !win.SetWindowPos(
			mw.hWnd, win.HWND_TOP,
			r.Left, r.Top, r.Right-r.Left, r.Bottom-r.Top,
			win.SWP_FRAMECHANGED|win.SWP_NOOWNERZORDER) {

			return lastError("SetWindowPos")
		}
	} else {
		if err := mw.ensureStyleBits(win.WS_OVERLAPPEDWINDOW, true); err != nil {
			return err
		}

		if !win.SetWindowPlacement(mw.hWnd, mw.windowPlacement) {
			return lastError("SetWindowPlacement")
		}

		if !win.SetWindowPos(mw.hWnd, 0, 0, 0, 0, 0, win.SWP_FRAMECHANGED|win.SWP_NOMOVE|
			win.SWP_NOOWNERZORDER|win.SWP_NOSIZE|win.SWP_NOZORDER) {

			return lastError("SetWindowPos")
		}
	}

	return nil
}

// WndProc returns a uintptr of the window process
func (mw *MainWindow) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_WINDOWPOSCHANGED, win.WM_SIZE:
		if win.WM_WINDOWPOSCHANGED == msg {
			wp := (*win.WINDOWPOS)(unsafe.Pointer(lParam))
			if wp.Flags&win.SWP_NOSIZE != 0 {
				break
			}
		}

		cb := mw.ClientBoundsPixels()

		if mw.toolBar != nil {
			bounds := Rectangle{0, 0, cb.Width, mw.toolBar.HeightPixels()}
			if mw.toolBar.BoundsPixels() != bounds {
				mw.toolBar.SetBoundsPixels(bounds)
			}
		}

		bounds := Rectangle{0, cb.Y + cb.Height, cb.Width, mw.statusBar.HeightPixels()}
		if mw.statusBar.BoundsPixels() != bounds {
			mw.statusBar.SetBoundsPixels(bounds)
		}
	}

	return mw.FormBase.WndProc(hwnd, msg, wParam, lParam)
}
