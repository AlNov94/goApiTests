package steps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertAll(t *testing.T) {
	passed := assert.False(t, t.Failed())
	if passed == false {
		t.FailNow()
	}
}

func SoftAssertContainsStep(description string, assert *assert.Assertions, actual interface{}, expected interface{}) {
	Step(description, func() bool {
		return assert.Contains(actual, expected)
	},
		"actual", actual,
		"expected", expected)
}

func SoftAssertEqualStep(description string, assert *assert.Assertions, actual interface{}, expected interface{}) {
	Step(description, func() bool {
		return assert.Equal(actual, expected)
	},
		"actual", actual,
		"expected", expected)
}

func SoftAssertEqualErrorStep(description string, assert *assert.Assertions, actual error, expected string) {
	Step(description, func() bool {
		return assert.EqualError(actual, expected)
	},
		"actual", actual,
		"expected", expected)
}

func SoftAssertEqualValuesStep(description string, assert *assert.Assertions, actual interface{}, expected interface{}) {
	Step(description, func() bool {
		return assert.EqualValues(actual, expected)
	},
		"actual", actual,
		"expected", expected)
}

func SoftAssertFalseStep(description string, assert *assert.Assertions, actual bool) {
	Step(description, func() bool {
		return assert.False(actual)
	},
		"actual", actual,
		"expected", false)
}

func SoftAssertTrueStep(description string, assert *assert.Assertions, actual bool) {
	Step(description, func() bool {
		return assert.True(actual)
	},
		"actual", actual,
		"expected", true)
}
