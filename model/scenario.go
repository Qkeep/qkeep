package model

type Scenario struct {
	GlobalSetting *GlobalSetting `yaml:"global_setting"`
}

func (s *Scenario) SetSetting(setting *GlobalSetting) {
	s.GlobalSetting = setting
}

func (s *Scenario) GetSetting() *GlobalSetting {
	return s.GlobalSetting
}
