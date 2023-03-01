package model

type Global struct {
	GlobalSetting *GlobalSetting `yaml:"global_setting"`
}

func (g *Global) SetSetting(setting *GlobalSetting) {
	g.GlobalSetting = setting
}

func (g *Global) GetSetting() *GlobalSetting {
	return g.GlobalSetting
}
