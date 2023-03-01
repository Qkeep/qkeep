package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobal(t *testing.T) {

	global := &Global{}

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

	// Test nil
	assert.Nil(t, global.GetSetting())

	global.SetSetting(globalSetting)

	assert := assert.New(t)

	// Can get the global setting.
	assert.Equal(globalSetting, global.GetSetting())

	// Can get the base URL.
	assert.Equal("http://localhost:8080", global.GetSetting().GetBaseUrl())

	// Can get the headers.
	headers := global.GetSetting().GetHeaders()
	assert.Equal(1, len(headers))
	assert.Equal("application/json", headers["Content-Type"])

	// Can get the flows.
	flows := global.GetSetting().GetFlows()
	assert.Equal(1, len(flows))
	assert.Equal("Flow 1", flows[0].Name)

	// Can get the flow by name.
	flow1 := global.GetSetting().GetFlowByName("Flow 1")
	assert.NotNil(flow1)

	// Can get the steps by flow name.
	steps1 := global.GetSetting().GetStepsByFlowName("Flow 1")
	assert.NotNil(steps1)
	assert.Equal(2, len(steps1))
}
