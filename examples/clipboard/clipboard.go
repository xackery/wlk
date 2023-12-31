// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/xackery/wlk/cpl"
	"github.com/xackery/wlk/walk"
)

func main() {
	var te *walk.TextEdit

	if _, err := (cpl.MainWindow{
		Title:   "Walk Clipboard Example",
		MinSize: cpl.Size{Width: 300, Height: 200},
		Layout:  cpl.VBox{},
		Children: []cpl.Widget{
			cpl.PushButton{
				Text: "Copy",
				OnClicked: func() {
					if err := walk.Clipboard().SetText(te.Text()); err != nil {
						log.Print("Copy: ", err)
					}
				},
			},
			cpl.PushButton{
				Text: "Paste",
				OnClicked: func() {
					if text, err := walk.Clipboard().Text(); err != nil {
						log.Print("Paste: ", err)
					} else {
						te.SetText(text)
					}
				},
			},
			cpl.TextEdit{
				AssignTo: &te,
			},
		},
	}).Run(); err != nil {
		log.Fatal(err)
	}
}
