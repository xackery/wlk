// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package walk

import (
	"github.com/xackery/wlk/win"
)

type EllipsisMode int

const (
	EllipsisNone EllipsisMode = 0
	EllipsisEnd               = EllipsisMode(win.SS_ENDELLIPSIS)
	EllipsisPath              = EllipsisMode(win.SS_PATHELLIPSIS)
)

type Label struct {
	static
	textChangedPublisher EventPublisher
}

func NewLabel(parent Container) (*Label, error) {
	return NewLabelWithStyle(parent, 0)
}

func NewLabelWithStyle(parent Container, style uint32) (*Label, error) {
	l := new(Label)

	if err := l.init(l, parent, style); err != nil {
		return nil, err
	}

	l.SetTextAlignment(AlignNear)

	l.MustRegisterProperty("Text", NewProperty(
		func() interface{} {
			return l.Text()
		},
		func(v interface{}) error {
			return l.SetText(assertStringOr(v, ""))
		},
		l.textChangedPublisher.Event()))

	return l, nil
}

func (l *Label) asStatic() *static {
	return &l.static
}

func (l *Label) EllipsisMode() EllipsisMode {
	style, err := win.GetWindowLong(l.hwndStatic, win.GWL_STYLE)
	if err != nil {
		//	fmt.Println("Error getting window long:", err)
	}

	return EllipsisMode(style & (win.SS_ENDELLIPSIS | win.SS_PATHELLIPSIS))
}

func (l *Label) SetEllipsisMode(mode EllipsisMode) error {
	oldMode := l.EllipsisMode()

	if mode == oldMode {
		return nil
	}

	if err := setAndClearWindowLongBits(l.hwndStatic, win.GWL_STYLE, uint32(mode), uint32(oldMode)); err != nil {
		return err
	}

	l.RequestLayout()

	return nil
}

func (l *Label) TextAlignment() Alignment1D {
	return l.textAlignment1D()
}

func (l *Label) SetTextAlignment(alignment Alignment1D) error {
	if alignment == AlignDefault {
		alignment = AlignNear
	}

	return l.setTextAlignment1D(alignment)
}

func (l *Label) Text() string {
	return l.text()
}

func (l *Label) SetText(text string) error {
	if changed, err := l.setText(text); err != nil {
		return err
	} else if !changed {
		return nil
	}

	l.textChangedPublisher.Publish()

	return nil
}
