package config

var loadedConfig Config

type Config struct {
	MemeLordApiUrl   string       `json:"memeLordApiUrl"`
	MemeLordApiToken string       `json:"memeLordApiToken"`
	Window           WindowConfig `json:"window"`
}

func GetConfig() Config {
	return loadedConfig
}
