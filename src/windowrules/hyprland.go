package windowrules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"meme-lord-picker/config"
	"os/exec"
)

type hyprland struct {
}

func (h hyprland) setWindowPositionRule() {
	cursorPosition, err := h.getCursorPosition()
	if err != nil {
		fmt.Println("Failed to create hyprland window position rule\n", err.Error())
		return
	}

	currentMonitor, err := h.getMonitorAtCursor(cursorPosition)
	if err != nil {
		fmt.Println("Failed to create hyprland window position rule\n", err.Error())
		return
	}

	targetPosition := h.adjustPositionBasedOnMonitor(cursorPosition, currentMonitor)

	h.createRule(targetPosition, currentMonitor)
}

func (h hyprland) getCursorPosition() (position, error) {
	cp := position{}
	output := bytes.NewBuffer(nil)
	cmd := exec.Command(`hyprctl`, "cursorpos", "-j")
	cmd.Stdout = output
	cmd.Stderr = output
	err := cmd.Run()
	if err != nil {
		return position{}, fmt.Errorf("failed to get cursorposition\n%s\n%s", err.Error(), output.String())
	}
	err = json.Unmarshal(output.Bytes(), &cp)
	if err != nil {
		return position{}, fmt.Errorf("failed to unmarshal cursorposition\n%s\n%s", err.Error(), output.String())
	}
	return cp, nil
}

func (h hyprland) getMonitorAtCursor(cp position) (monitor, error) {
	monitors := []monitor{}

	cmd := exec.Command("hyprctl", "monitors", "-j")
	output := bytes.NewBuffer(nil)
	cmd.Stdout = output
	cmd.Stderr = output
	err := cmd.Run()
	if err != nil {
		return monitor{}, fmt.Errorf("failed to get monitors\n%s\n%s", err.Error(), output.String())
	}

	err = json.Unmarshal(output.Bytes(), &monitors)
	if err != nil {
		return monitor{}, fmt.Errorf("failed to unmarshal monitors\n%s\n%s", err.Error(), output.String())
	}

	for _, monitor := range monitors {
		xMax := monitor.X + monitor.Width
		yMax := monitor.Y + monitor.Height
		if monitor.X <= cp.X && cp.X <= xMax && monitor.Y <= cp.Y && cp.Y <= yMax {
			return monitor, nil
		}
	}

	return monitors[0], nil
}

func (h hyprland) adjustPositionBasedOnMonitor(cursorPos position, mon monitor) position {
	conf := config.GetWindowConfig()
	cursorPos.X -= mon.X
	cursorPos.Y -= mon.Y
	result := cursorPos
	maxX := mon.Width
	maxY := mon.Height
	if result.X+conf.Width > maxX {
		result.X = result.X - conf.Width
	}
	if result.Y+conf.Height > maxY {
		result.Y = result.Y - conf.Height
	}
	return result
}

func (h hyprland) createRule(pos position, mon monitor) error {
	output := bytes.NewBuffer(nil)
	ruleCommand := fmt.Sprintf(`hl.window_rule({
		name = "MemeLord Picker window position temp rule",
		match = {
			class = "^%s$"
		},
		move = {%d, %d},
		monitor = %d
	})`, _APP_CLASS, pos.X, pos.Y, mon.ID)

	cmd := exec.Command("hyprctl", "eval", ruleCommand)
	cmd.Stderr = output
	cmd.Stdout = output
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create hyprland window position rule\n%s\n%s", err.Error(), output.String())
	}
	return nil
}
