package model

type Suit struct {
	GlobalSetting *GlobalSetting `yaml:"global_setting"`
}

func (s *Suit) SetSetting(setting *GlobalSetting) {
	s.GlobalSetting = setting
}

func (s *Suit) GetSetting() *GlobalSetting {
	return s.GlobalSetting
}
