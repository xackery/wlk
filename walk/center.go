package walk

import "github.com/xackery/wlk/win"

// CenterWindowOnScreen will center the window on the screen
func CenterWindowOnScreen(window *MainWindow) {
	// Get the screen dimensions
	screenWidth := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	screenHeight := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	// Get the window dimensions
	windowWidth := int(window.Width())
	windowHeight := int(window.Height())

	// Calculate the new window position
	x := (screenWidth - windowWidth) / 2
	y := (screenHeight - windowHeight) / 2

	// Move the window to the center of the screen
	window.SetBounds(Rectangle{X: x, Y: y, Width: windowWidth, Height: windowHeight})
}
