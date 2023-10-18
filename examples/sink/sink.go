// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/xackery/wlk/cpl"
	"github.com/xackery/wlk/walk"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("Failed to run:", err)
		os.Exit(1)
	}
}

func run() error {

	var mw *walk.MainWindow
	cmw := cpl.MainWindow{
		Title:    "Kitchen Sink Example",
		Name:     "sink",
		AssignTo: &mw,
		Size:     cpl.Size{Width: 365, Height: 371},
		Layout:   cpl.VBox{},
		Children: []cpl.Widget{
			cpl.Label{Text: "Label"},
			cpl.LineEdit{Text: "LineEdit"},
			cpl.TextEdit{Text: "TextEdit"},
			cpl.CheckBox{Text: "CheckBox"},
			cpl.RadioButton{Text: "RadioButton"},
			cpl.GroupBox{
				Title:  "GroupBox",
				Layout: cpl.VBox{},
				Children: []cpl.Widget{
					cpl.RadioButton{Text: "RadioButton"},
					cpl.RadioButton{Text: "RadioButton"},
					cpl.RadioButton{Text: "RadioButton"},
				},
			},
			cpl.ComboBox{
				Value:    "ComboBox",
				Editable: true,
				Model:    []string{"Item1", "Item2", "Item3"},
			},
			// composite
			// customwidget
			// databinder
			cpl.DateEdit{Date: time.Now()},
			cpl.DateLabel{Date: time.Now()},
			cpl.PushButton{
				Text: "Dialog",
				OnClicked: func() {
					walk.MsgBox(mw, "Message", "Message", walk.MsgBoxIconInformation)
				},
			},
			// gradientcomposite
			cpl.ImageView{},
			cpl.LinkLabel{Text: "LinkLabel"},
			cpl.ListBox{
				Model: []string{"Item1", "Item2", "Item3"},
			},
			cpl.NumberEdit{
				Value: 1,
			},
			cpl.NumberLabel{
				Value: 1,
			},
			cpl.ProgressBar{
				Value: 50,
			},
			// radiobuttongroup
			// radiobuttongroupbox
			cpl.ScrollView{
				Layout: cpl.VBox{},
				Children: []cpl.Widget{
					cpl.Label{Text: "Label"},
					cpl.LineEdit{Text: "LineEdit"},
					cpl.TextEdit{Text: "TextEdit"},
				},
			},
			// seperator
			cpl.Slider{
				Value: 50,
			},
			// spacer
			cpl.SplitButton{
				Text: "SplitButton",
			},
			// splitter
			/*cpl.TableView{
				Columns: []cpl.TableViewColumn{
					{Title: "Column1"},
					{Title: "Column2"},
					{Title: "Column3"},
				},
				Model: [][]interface{}{
					{"Item1", "Item2", "Item3"},
					{"Item1", "Item2", "Item3"},
					{"Item1", "Item2", "Item3"},
				},
			},*/
			cpl.TabPage{
				Title:  "TabPage",
				Layout: cpl.VBox{},
				Children: []cpl.Widget{
					cpl.Label{Text: "Label"},
					cpl.LineEdit{Text: "LineEdit"},
					cpl.TextEdit{Text: "TextEdit"},
				},
			},
			cpl.TabWidget{
				//Layout: cpl.VBox{},
				Pages: []cpl.TabPage{
					{
						Title:  "TabPage",
						Layout: cpl.VBox{},
						Children: []cpl.Widget{
							cpl.Label{Text: "Label"},
							cpl.LineEdit{Text: "LineEdit"},
							cpl.TextEdit{Text: "TextEdit"},
						},
					},
					{
						Title:  "TabPage",
						Layout: cpl.VBox{},
						Children: []cpl.Widget{
							cpl.Label{Text: "Label"},
							cpl.LineEdit{Text: "LineEdit"},
							cpl.TextEdit{Text: "TextEdit"},
						},
					},
				},
			},
			cpl.TextLabel{Text: "TextLabel"},
			/*cpl.TreeView{
				Model: []cpl.TreeItem{
					cpl.TreeNode{
						Text: "TreeNode",
						Items: []cpl.TreeItem{
							cpl.TreeNode{
								Text: "TreeNode",
								Items: []cpl.TreeItem{
									cpl.TreeNode{
										Text: "TreeNode",
									},
								},
							},
						},
					},
				},
			},*/
			cpl.WebView{
				URL: "https://github.com",
			},
		},
		Visible: false,
	}

	err := cmw.Create()
	if err != nil {
		return fmt.Errorf("create main window: %w", err)
	}

	mw.SetVisible(true)
	code := mw.Run()
	if code != 0 {
		return fmt.Errorf("exit code %d", code)
	}
	return nil
}
