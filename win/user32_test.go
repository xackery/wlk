package win

import (
	"fmt"
	"testing"
	"time"
)

func TestGetActiveWindowTitle(t *testing.T) {
	title := GetActiveWindowTitle()
	fmt.Println("title:", title)
}

func TestKeyPress(t *testing.T) {
	//KeyPress(VK_CONTROL, VK_MENU, VK_DELETE)
	KeyPress(VK_1, 10*time.Millisecond)
}
