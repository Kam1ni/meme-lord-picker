package config

type WindowConfig struct {
	Width       int
	Height      int
	ImageWidth  int
	ImageHeight int
}

func GetWindowConfig() WindowConfig {
	return loadedConfig.Window
}

func (w *WindowConfig) enforceLimits() {
	if w.Width < 10 {
		w.Width = 10
	}
	if w.Height < 10 {
		w.Height = 10
	}
	if w.ImageWidth < 10 {
		w.ImageWidth = 10
	}
	if w.ImageHeight < 10 {
		w.ImageHeight = 10
	}
}
