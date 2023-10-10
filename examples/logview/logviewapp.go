// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"time"

	"github.com/xackery/wlk/cpl"
	"github.com/xackery/wlk/walk"
)

func main() {
	var mw *walk.MainWindow

	if err := (cpl.MainWindow{
		AssignTo: &mw,
		Title:    "Walk LogView Example",
		MinSize:  cpl.Size{Width: 320, Height: 240},
		Size:     cpl.Size{Width: 400, Height: 600},
		Layout:   cpl.VBox{MarginsZero: true},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	lv, err := NewLogView(mw)
	if err != nil {
		log.Fatal(err)
	}

	lv.PostAppendText("XXX")
	log.SetOutput(lv)

	go func() {
		for i := 0; i < 10000; i++ {
			time.Sleep(100 * time.Millisecond)
			log.Println("Text")
		}
	}()

	mw.Run()
}
