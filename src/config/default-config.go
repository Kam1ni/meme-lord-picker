package config

func getDefaultConfig() Config {
	return Config{
		MemeLordApiUrl:   "",
		MemeLordApiToken: "",
		Window: WindowConfig{
			Width:       400,
			Height:      400,
			ImageWidth:  200,
			ImageHeight: 200,
		},
	}

}
