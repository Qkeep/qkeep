package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalSetting(t *testing.T) {
	// Create a new GlobalSetting instance.
	g := &GlobalSetting{
		BaseUrl: "http://example.com",
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
			{
				Name: "Flow 2",
				Steps: []string{
					"Step 3",
					"Step 4",
				},
			},
		},
	}

	// Test GetBaseUrl method.
	assert.Equal(t, "http://example.com", g.GetBaseUrl())

	// Test GetHeaders method.
	headers := g.GetHeaders()
	assert.Equal(t, 1, len(headers))
	assert.Equal(t, "application/json", headers["Content-Type"])

	// Test GetFlows method.
	flows := g.GetFlows()
	assert.Equal(t, 2, len(flows))
	assert.Equal(t, "Flow 1", flows[0].Name)
	assert.Equal(t, "Flow 2", flows[1].Name)

	// Test GetFlowByName method.
	flow1 := g.GetFlowByName("Flow 1")
	assert.NotNil(t, flow1)
	assert.Equal(t, "Flow 1", flow1.Name)
	assert.Equal(t, 2, len(flow1.Steps))

	flow3 := g.GetFlowByName("Flow 3")
	assert.Nil(t, flow3)

	// Test GetStepsByFlowName method.
	steps1 := g.GetStepsByFlowName("Flow 1")
	assert.NotNil(t, steps1)
	assert.Equal(t, 2, len(steps1))
	assert.Equal(t, "Step 1", steps1[0])
	assert.Equal(t, "Step 2", steps1[1])

	steps3 := g.GetStepsByFlowName("Flow 3")
	assert.Nil(t, steps3)

	// Test Merge method.
	other := &GlobalSetting{
		BaseUrl: "http://example.org",
		Headers: map[string]string{
			"Content-Type": "application/xml",
		},
		Flows: []Flow{
			{
				Name: "Flow 3",
				Steps: []string{
					"Step 5",
					"Step 6",
				},
			},
		},
	}

	g.Merge(other)
	assert.Equal(t, "http://example.org", g.BaseUrl)
	assert.Equal(t, 1, len(g.Headers))
	assert.Equal(t, "application/xml", g.Headers["Content-Type"])
	assert.Equal(t, 1, len(g.Flows))
	assert.Equal(t, "Flow 3", g.Flows[0].Name)
}
