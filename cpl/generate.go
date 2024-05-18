package cpl

func GenerateComposite(layout Layout, widgets ...Widget) Composite {
	return Composite{
		Layout:   layout,
		Children: widgets,
	}
}
