package baseSteps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//testify hard assert wrapper for allure reporting

func AssertContainsStep(description string, t *testing.T, actual interface{}, expected interface{}) {
	StepNoResult(description, func() {
		assert.Contains(t, actual, expected)
		AssertAll(t)
	},
		Parameter{"actual", actual},
		Parameter{"expected", expected})
}

func AssertEqualStep(description string, t *testing.T, actual interface{}, expected interface{}) {
	StepNoResult(description, func() {
		assert.Equal(t, actual, expected)
		AssertAll(t)
	},
		Parameter{"actual", actual},
		Parameter{"expected", expected})
}

func AssertEqualErrorStep(description string, t *testing.T, actual error, expected string) {
	StepNoResult(description, func() {
		assert.EqualError(t, actual, expected)
		AssertAll(t)
	},
		Parameter{"actual", actual},
		Parameter{"expected", expected})
}

func AssertEqualValuesStep(description string, t *testing.T, actual interface{}, expected interface{}) {
	StepNoResult(description, func() {
		assert.EqualValues(t, actual, expected)
		AssertAll(t)
	},
		Parameter{"actual", actual},
		Parameter{"expected", expected})
}

func AssertFalseStep(description string, t *testing.T, actual bool) {
	StepNoResult(description, func() {
		assert.False(t, actual)
		AssertAll(t)
	},
		Parameter{"actual", actual},
		Parameter{"expected", false})
}

func AssertTrueStep(description string, t *testing.T, actual bool) {
	StepNoResult(description, func() {
		assert.True(t, actual)
		AssertAll(t)
	},
		Parameter{"actual", actual},
		Parameter{"expected", true})
}
