// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/xackery/wlk/cpl"
	"github.com/xackery/wlk/walk"
)

var isSpecialMode = walk.NewMutableCondition()

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {
	cpl.MustRegisterCondition("isSpecialMode", isSpecialMode)

	mw := new(MyMainWindow)

	var openAction, showAboutBoxAction *walk.Action
	var recentMenu *walk.Menu
	var toggleSpecialModePB *walk.PushButton

	if err := (cpl.MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Walk Actions Example",
		MenuItems: []cpl.MenuItem{
			cpl.Menu{
				Text: "&File",
				Items: []cpl.MenuItem{
					cpl.Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						Image:       "../img/open.png",
						Enabled:     cpl.Bind("enabledCB.Checked"),
						Visible:     cpl.Bind("!openHiddenCB.Checked"),
						Shortcut:    cpl.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO},
						OnTriggered: mw.openAction_Triggered,
					},
					cpl.Menu{
						AssignTo: &recentMenu,
						Text:     "Recent",
					},
					cpl.Separator{},
					cpl.Action{
						Text:        "E&xit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			cpl.Menu{
				Text: "&View",
				Items: []cpl.MenuItem{
					cpl.Action{
						Text:    "Open / Special Enabled",
						Checked: cpl.Bind("enabledCB.Visible"),
					},
					cpl.Action{
						Text:    "Open Hidden",
						Checked: cpl.Bind("openHiddenCB.Visible"),
					},
				},
			},
			cpl.Menu{
				Text: "&Help",
				Items: []cpl.MenuItem{
					cpl.Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "About",
						OnTriggered: mw.showAboutBoxAction_Triggered,
					},
				},
			},
		},
		ToolBar: cpl.ToolBar{
			ButtonStyle: cpl.ToolBarButtonImageBeforeText,
			Items: []cpl.MenuItem{
				cpl.ActionRef{Action: &openAction},
				cpl.Menu{
					Text:  "New A",
					Image: "../img/document-new.png",
					Items: []cpl.MenuItem{
						cpl.Action{
							Text:        "A",
							OnTriggered: mw.newAction_Triggered,
						},
						cpl.Action{
							Text:        "B",
							OnTriggered: mw.newAction_Triggered,
						},
						cpl.Action{
							Text:        "C",
							OnTriggered: mw.newAction_Triggered,
						},
					},
					OnTriggered: mw.newAction_Triggered,
				},
				cpl.Separator{},
				cpl.Menu{
					Text:  "View",
					Image: "../img/document-properties.png",
					Items: []cpl.MenuItem{
						cpl.Action{
							Text:        "X",
							OnTriggered: mw.changeViewAction_Triggered,
						},
						cpl.Action{
							Text:        "Y",
							OnTriggered: mw.changeViewAction_Triggered,
						},
						cpl.Action{
							Text:        "Z",
							OnTriggered: mw.changeViewAction_Triggered,
						},
					},
				},
				cpl.Separator{},
				cpl.Action{
					Text:        "Special",
					Image:       "../img/system-shutdown.png",
					Enabled:     cpl.Bind("isSpecialMode && enabledCB.Checked"),
					OnTriggered: mw.specialAction_Triggered,
				},
			},
		},
		ContextMenuItems: []cpl.MenuItem{
			cpl.ActionRef{Action: &showAboutBoxAction},
		},
		MinSize: cpl.Size{Width: 300, Height: 200},
		Layout:  cpl.VBox{},
		Children: []cpl.Widget{
			cpl.CheckBox{
				Name:    "enabledCB",
				Text:    "Open / Special Enabled",
				Checked: true,
				Accessibility: cpl.Accessibility{
					Help: "Enables Open and Special",
				},
			},
			cpl.CheckBox{
				Name:    "openHiddenCB",
				Text:    "Open Hidden",
				Checked: true,
			},
			cpl.PushButton{
				AssignTo: &toggleSpecialModePB,
				Text:     "Enable Special Mode",
				OnClicked: func() {
					isSpecialMode.SetSatisfied(!isSpecialMode.Satisfied())

					if isSpecialMode.Satisfied() {
						toggleSpecialModePB.SetText("Disable Special Mode")
					} else {
						toggleSpecialModePB.SetText("Enable Special Mode")
					}
				},
				Accessibility: cpl.Accessibility{
					Help: "Toggles special mode",
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	addRecentFileActions := func(texts ...string) {
		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(mw.openAction_Triggered)
			recentMenu.Actions().Add(a)
		}
	}

	addRecentFileActions("Foo", "Bar", "Baz")

	mw.Run()
}

func (mw *MyMainWindow) openAction_Triggered() {
	walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) newAction_Triggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) changeViewAction_Triggered() {
	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
	walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) specialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
