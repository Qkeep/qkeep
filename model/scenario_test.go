package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScenario(t *testing.T) {

	globalSetting := &GlobalSetting{
		BaseUrl: "http://localhost:8080",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Flows: []Flow{
			{
				Name: "Flow 1",
				Steps: []string{
					"Step 1",
					"Step 2",
				},
			},
		},
	}
	scenario := &Scenario{}

	// Test nil
	assert.Nil(t, scenario.GetSetting())

	scenario.SetSetting(globalSetting)

	assert := assert.New(t)

	// Can get the global setting.
	assert.Equal(globalSetting, scenario.GetSetting())

	// Can get the base URL.
	assert.Equal("http://localhost:8080", scenario.GetSetting().GetBaseUrl())

	// Can get the headers.
	headers := scenario.GetSetting().GetHeaders()
	assert.Equal(1, len(headers))
	assert.Equal("application/json", headers["Content-Type"])

	// Can get the flows.
	flows := scenario.GetSetting().GetFlows()
	assert.Equal(1, len(flows))
	assert.Equal("Flow 1", flows[0].Name)

	// Can get the flow by name.
	flow1 := scenario.GetSetting().GetFlowByName("Flow 1")
	assert.NotNil(flow1)

	// Can get the steps by flow name.
	steps1 := scenario.GetSetting().GetStepsByFlowName("Flow 1")
	assert.NotNil(steps1)
	assert.Equal(2, len(steps1))
}
