package model

type GlobalSetting struct {
	BaseUrl string            `yaml:"base_url"`
	Headers map[string]string `yaml:"headers"`
	Flows   []Flow            `yaml:"flows"`
}

type Flow struct {
	Name  string   `yaml:"name"`
	Steps []string `yaml:"steps"`
}

// Return the base URL.
func (g *GlobalSetting) GetBaseUrl() string {
	return g.BaseUrl
}

// Return the headers.
func (g *GlobalSetting) GetHeaders() map[string]string {
	return g.Headers
}

// Return the flows.
func (g *GlobalSetting) GetFlows() []Flow {
	return g.Flows
}

// Return the flow by name.
func (g *GlobalSetting) GetFlowByName(name string) *Flow {
	for _, flow := range g.Flows {
		if flow.Name == name {
			return &flow
		}
	}
	return nil
}

// Return the steps by flow name.
func (g *GlobalSetting) GetStepsByFlowName(name string) []string {
	for _, flow := range g.Flows {
		if flow.Name == name {
			return flow.Steps
		}
	}
	return nil
}

// Merge the other global setting.
// If the other global setting is nil, return the current global setting.
// If the other global setting is not nil, merge (override) the other global setting to the current global setting.
func (g *GlobalSetting) Merge(other *GlobalSetting) *GlobalSetting {
	if other == nil {
		return g
	}
	if other.BaseUrl != "" {
		g.BaseUrl = other.BaseUrl
	}
	if other.Headers != nil {
		g.Headers = other.Headers
	}
	if other.Flows != nil {
		g.Flows = other.Flows
	}
	return g
}
