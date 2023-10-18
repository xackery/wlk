// Copyright 2018 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package cpl

import (
	"github.com/xackery/wlk/common"
	"github.com/xackery/wlk/walk"
)

type DateLabel struct {
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

	// static

	TextColor common.Color

	// DateLabel

	AssignTo      **walk.DateLabel
	Date          Property
	Format        Property
	TextAlignment Alignment1D
}

func (dl DateLabel) Create(builder *Builder) error {
	w, err := walk.NewDateLabel(builder.Parent())
	if err != nil {
		return err
	}

	if dl.AssignTo != nil {
		*dl.AssignTo = w
	}

	return builder.InitWidget(dl, w, func() error {
		if err := w.SetTextAlignment(walk.Alignment1D(dl.TextAlignment)); err != nil {
			return err
		}

		w.SetTextColor(dl.TextColor)

		return nil
	})
}
