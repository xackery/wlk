// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package cpl

import (
	"fmt"

	"github.com/xackery/wlk/common"
	"github.com/xackery/wlk/walk"
	"github.com/xackery/wlk/wcolor"
	"github.com/xackery/wlk/win"
)

type TextEdit struct {
	// Window

	Accessibility      Accessibility
	Background         Brush
	ContextMenuItems   []MenuItem
	DoubleBuffering    bool
	Enabled            Property
	Font               Font
	MaxSize            Size
	MinSize            Size
	Name               string
	OnBoundsChanged    walk.EventHandler
	OnKeyDown          walk.KeyEventHandler
	OnKeyPress         walk.KeyEventHandler
	OnKeyUp            walk.KeyEventHandler
	OnMouseDown        walk.MouseEventHandler
	OnMouseMove        walk.MouseEventHandler
	OnMouseUp          walk.MouseEventHandler
	OnSizeChanged      walk.EventHandler
	Persistent         bool
	RightToLeftReading bool
	ToolTipText        Property
	Visible            Property

	// Widget

	Alignment          Alignment2D
	AlwaysConsumeSpace bool
	Column             int
	ColumnSpan         int
	GraphicsEffects    []walk.WidgetGraphicsEffect
	Row                int
	RowSpan            int
	StretchFactor      int

	// TextEdit

	AssignTo      **walk.TextEdit
	CompactHeight bool
	HScroll       bool
	MaxLength     int
	OnTextChanged walk.EventHandler
	ReadOnly      Property
	Text          Property
	TextAlignment Alignment1D
	TextColor     wcolor.Color
	VScroll       bool
}

func (te TextEdit) Create(builder *Builder) error {
	var style uint32
	if te.HScroll {
		style |= win.WS_HSCROLL
	}
	if te.VScroll {
		style |= win.WS_VSCROLL
	}

	w, err := walk.NewTextEditWithStyle(builder.Parent(), style)
	if err != nil {
		return err
	}

	if te.AssignTo != nil {
		*te.AssignTo = w
	}

	return builder.InitWidget(te, w, func() error {
		w.SetCompactHeight(te.CompactHeight)
		if IsDarkMode() {
			if te.TextColor == 0 {
				te.TextColor = common.DarkTextFG
			}

			te.Background = SolidColorBrush{Color: common.DarkFormLighterBG}
			brush, err := walk.NewSolidColorBrush(common.DarkFormLighterBG)
			if err != nil {
				return fmt.Errorf("new solid color brush: %w", err)
			}
			w.SetBackground(brush)
		}

		w.SetTextColor(te.TextColor)

		if err := w.SetTextAlignment(walk.Alignment1D(te.TextAlignment)); err != nil {
			return err
		}

		if te.MaxLength > 0 {
			w.SetMaxLength(te.MaxLength)
		}

		if te.OnTextChanged != nil {
			w.TextChanged().Attach(te.OnTextChanged)
		}

		return nil
	})
}
