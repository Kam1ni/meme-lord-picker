package config

type WindowConfig struct {
	Width       int
	Height      int
	ImageWidth  int
	ImageHeight int
}

func GetWindowConfig() WindowConfig {
	return WindowConfig{
		Width:       400,
		Height:      400,
		ImageWidth:  200,
		ImageHeight: 200,
	}
}
