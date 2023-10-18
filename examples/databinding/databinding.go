// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xackery/wlk/common"
	"github.com/xackery/wlk/walk"

	"github.com/xackery/wlk/cpl"
)

func main() {
	walk.AppendToWalkInit(func() {
		walk.FocusEffect, _ = walk.NewBorderGlowEffect(common.RGB(0, 63, 255))
		walk.InteractionEffect, _ = walk.NewDropShadowEffect(common.RGB(63, 63, 63))
		walk.ValidationErrorEffect, _ = walk.NewBorderGlowEffect(common.RGB(255, 0, 0))
	})

	var mw *walk.MainWindow
	var outTE *walk.TextEdit

	animal := new(Animal)

	if _, err := (cpl.MainWindow{
		AssignTo: &mw,
		Title:    "Walk Data Binding Example",
		MinSize:  cpl.Size{Width: 300, Height: 200},
		Layout:   cpl.VBox{},
		Children: []cpl.Widget{
			cpl.PushButton{
				Text: "Edit Animal",
				OnClicked: func() {
					if cmd, err := RunAnimalDialog(mw, animal); err != nil {
						log.Print(err)
					} else if cmd == walk.DlgCmdOK {
						outTE.SetText(fmt.Sprintf("%+v", animal))
					}
				},
			},
			cpl.Label{
				Text: "animal:",
			},
			cpl.TextEdit{
				AssignTo: &outTE,
				ReadOnly: true,
				Text:     fmt.Sprintf("%+v", animal),
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

type Animal struct {
	Name          string
	ArrivalDate   time.Time
	SpeciesId     int
	Speed         int
	Sex           Sex
	Weight        float64
	PreferredFood string
	Domesticated  bool
	Remarks       string
	Patience      time.Duration
}

func (a *Animal) PatienceField() *DurationField {
	return &DurationField{&a.Patience}
}

type Species struct {
	Id   int
	Name string
}

func KnownSpecies() []*Species {
	return []*Species{
		{1, "Dog"},
		{2, "Cat"},
		{3, "Bird"},
		{4, "Fish"},
		{5, "Elephant"},
	}
}

type DurationField struct {
	p *time.Duration
}

func (*DurationField) CanSet() bool       { return true }
func (f *DurationField) Get() interface{} { return f.p.String() }
func (f *DurationField) Set(v interface{}) error {
	x, err := time.ParseDuration(v.(string))
	if err == nil {
		*f.p = x
	}
	return err
}
func (f *DurationField) Zero() interface{} { return "" }

type Sex byte

const (
	SexMale Sex = 1 + iota
	SexFemale
	SexHermaphrodite
)

func RunAnimalDialog(owner walk.Form, animal *Animal) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return cpl.Dialog{
		AssignTo:      &dlg,
		Title:         cpl.Bind("'Animal Details' + (animal.Name == '' ? '' : ' - ' + animal.Name)"),
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: cpl.DataBinder{
			AssignTo:       &db,
			Name:           "animal",
			DataSource:     animal,
			ErrorPresenter: cpl.ToolTipErrorPresenter{},
		},
		MinSize: cpl.Size{300, 300},
		Layout:  cpl.VBox{},
		Children: []cpl.Widget{
			cpl.Composite{
				Layout: cpl.Grid{Columns: 2},
				Children: []cpl.Widget{
					cpl.Label{
						Text: "Name:",
					},
					cpl.LineEdit{
						Text: cpl.Bind("Name"),
					},

					cpl.Label{
						Text: "Arrival Date:",
					},
					cpl.DateEdit{
						Date: cpl.Bind("ArrivalDate"),
					},

					cpl.Label{
						Text: "Species:",
					},
					cpl.ComboBox{
						Value:         cpl.Bind("SpeciesId", cpl.SelRequired{}),
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         KnownSpecies(),
					},

					cpl.Label{
						Text: "Speed:",
					},
					cpl.Slider{
						Value: cpl.Bind("Speed"),
					},

					cpl.RadioButtonGroupBox{
						ColumnSpan: 2,
						Title:      "Sex",
						Layout:     cpl.HBox{},
						DataMember: "Sex",
						Buttons: []cpl.RadioButton{
							{Text: "Male", Value: SexMale},
							{Text: "Female", Value: SexFemale},
							{Text: "Hermaphrodite", Value: SexHermaphrodite},
						},
					},

					cpl.Label{
						Text: "Weight:",
					},
					cpl.NumberEdit{
						Value:    cpl.Bind("Weight", cpl.Range{0.01, 9999.99}),
						Suffix:   " kg",
						Decimals: 2,
					},

					cpl.Label{
						Text: "Preferred Food:",
					},
					cpl.ComboBox{
						Editable: true,
						Value:    cpl.Bind("PreferredFood"),
						Model:    []string{"Fruit", "Grass", "Fish", "Meat"},
					},

					cpl.Label{
						Text: "Domesticated:",
					},
					cpl.CheckBox{
						Checked: cpl.Bind("Domesticated"),
					},

					cpl.VSpacer{
						ColumnSpan: 2,
						Size:       8,
					},

					cpl.Label{
						ColumnSpan: 2,
						Text:       "Remarks:",
					},
					cpl.TextEdit{
						ColumnSpan: 2,
						MinSize:    cpl.Size{100, 50},
						Text:       cpl.Bind("Remarks"),
					},

					cpl.Label{
						ColumnSpan: 2,
						Text:       "Patience:",
					},
					cpl.LineEdit{
						ColumnSpan: 2,
						Text:       cpl.Bind("PatienceField"),
					},
				},
			},
			cpl.Composite{
				Layout: cpl.HBox{},
				Children: []cpl.Widget{
					cpl.HSpacer{},
					cpl.PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()
						},
					},
					cpl.PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}
