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
	"github.com/xackery/wlk/win"
)

type ListBox struct {
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

	// ListBox

	AssignTo                 **walk.ListBox
	BindingMember            string
	CurrentIndex             Property
	DisplayMember            string
	Format                   string
	ItemStyler               walk.ListItemStyler
	Model                    interface{}
	MultiSelection           bool
	OnCurrentIndexChanged    walk.EventHandler
	OnItemActivated          walk.EventHandler
	OnSelectedIndexesChanged walk.EventHandler
	Precision                int
	Value                    Property
}

func (lb ListBox) Create(builder *Builder) error {
	var w *walk.ListBox
	var err error
	if _, ok := lb.Model.([]string); ok &&
		(lb.BindingMember != "" || lb.DisplayMember != "") {

		return fmt.Errorf("listBox.Create: BindingMember and DisplayMember must be empty for []string models")
	}

	var style uint32

	if lb.ItemStyler != nil {
		style |= win.LBS_OWNERDRAWVARIABLE
	}
	if lb.MultiSelection {
		style |= win.LBS_EXTENDEDSEL
	}

	w, err = walk.NewListBoxWithStyle(builder.Parent(), style)
	if err != nil {
		return err
	}

	if lb.AssignTo != nil {
		*lb.AssignTo = w
	}

	return builder.InitWidget(lb, w, func() error {
		if lb.ItemStyler != nil {
			w.SetItemStyler(lb.ItemStyler)
		}
		w.SetFormat(lb.Format)
		w.SetPrecision(lb.Precision)

		if IsDarkMode() {
			brush, err := walk.NewSolidColorBrush(common.DarkFormLighterBG)
			if err != nil {
				return fmt.Errorf("new solid color brush: %w", err)
			}
			w.SetBackground(brush)
		}

		if err := w.SetBindingMember(lb.BindingMember); err != nil {
			return err
		}
		if err := w.SetDisplayMember(lb.DisplayMember); err != nil {
			return err
		}

		if err := w.SetModel(lb.Model); err != nil {
			return err
		}

		if lb.OnCurrentIndexChanged != nil {
			w.CurrentIndexChanged().Attach(lb.OnCurrentIndexChanged)
		}
		if lb.OnSelectedIndexesChanged != nil {
			w.SelectedIndexesChanged().Attach(lb.OnSelectedIndexesChanged)
		}
		if lb.OnItemActivated != nil {
			w.ItemActivated().Attach(lb.OnItemActivated)
		}

		return nil
	})
}
