package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuit(t *testing.T) {
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

	suit := &Suit{}

	// Test nil
	assert.Nil(t, suit.GetSetting())

	suit.SetSetting(globalSetting)

	assert := assert.New(t)

	// Can get the global setting.
	assert.Equal(globalSetting, suit.GetSetting())

	// Can get the base URL.
	assert.Equal("http://localhost:8080", suit.GetSetting().GetBaseUrl())

	// Can get the headers.
	headers := suit.GetSetting().GetHeaders()
	assert.Equal(1, len(headers))
	assert.Equal("application/json", headers["Content-Type"])

	// Can get the flows.
	flows := suit.GetSetting().GetFlows()
	assert.Equal(1, len(flows))
	assert.Equal("Flow 1", flows[0].Name)

	// Can get the flow by name.
	flow1 := suit.GetSetting().GetFlowByName("Flow 1")
	assert.NotNil(flow1)

	// Can get the steps by flow name.
	steps1 := suit.GetSetting().GetStepsByFlowName("Flow 1")
	assert.NotNil(steps1)
	assert.Equal(2, len(steps1))

}
