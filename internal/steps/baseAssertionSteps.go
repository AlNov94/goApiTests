package steps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertContainsStep(description string, t *testing.T, actual interface{}, expected interface{}) {
	Step(description, func() bool {
		return assert.Contains(t, actual, expected)
	},
		"actual", actual,
		"expected", expected)
	AssertAll(t)
}

func AssertEqualStep(description string, t *testing.T, actual interface{}, expected interface{}) {
	Step(description, func() bool {
		return assert.Equal(t, actual, expected)
	},
		"actual", actual,
		"expected", expected)
	AssertAll(t)
}

func AssertEqualErrorStep(description string, t *testing.T, actual error, expected string) {
	Step(description, func() bool {
		return assert.EqualError(t, actual, expected)
	},
		"actual", actual,
		"expected", expected)
	AssertAll(t)
}

func AssertEqualValuesStep(description string, t *testing.T, actual interface{}, expected interface{}) {
	Step(description, func() bool {
		return assert.EqualValues(t, actual, expected)
	},
		"actual", actual,
		"expected", expected)
	AssertAll(t)
}

func AssertFalseStep(description string, t *testing.T, actual bool) {
	Step(description, func() bool {
		return assert.False(t, actual)
	},
		"actual", actual,
		"expected", false)
	AssertAll(t)
}

func AssertTrueStep(description string, t *testing.T, actual bool) {
	Step(description, func() bool {
		return assert.True(t, actual)
	},
		"actual", actual,
		"expected", true)
	AssertAll(t)
}
