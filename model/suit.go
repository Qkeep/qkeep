package model

type Suit struct {
	Name      string                 `yaml:"name"`
	BaseUrl   string                 `yaml:"base_url"`
	Headers   map[string]interface{} `yaml:"headers"`
	Path      string                 `yaml:"path"`
	Version   string                 `yaml:"version"`
	Flow      []string               `yaml:"flow"`
	Variables map[string]interface{} `yaml:"variables"`
	Models    map[string]*Model
}
