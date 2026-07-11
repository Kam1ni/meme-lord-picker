package windowrules

import (
	"fmt"
	"os"
)

const _APP_CLASS = "meme-lord-picker"

type monitor struct {
	ID     int `json:"id"`
	Width  int `json:"width"`
	Height int `json:"height"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

type position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func AttemptSetWindowPositionRule() {
	desktop := os.Getenv("XDG_SESSION_DESKTOP")
	de, ok := desktopEnvironments[desktop]
	if !ok {
		fmt.Printf("No window positioning handler defined for %s.\n", desktop)
		return
	}
	de.setWindowPositionRule()
}
