package steps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertContainsStep(description string, t *testing.T, actual interface{}, expected interface{}) {
	SoftAssertContainsStep(description, assert.New(t), actual, expected)
}

func AssertEqualStep(description string, t *testing.T, actual interface{}, expected interface{}) {
	SoftAssertEqualStep(description, assert.New(t), actual, expected)
}

func AssertEqualErrorStep(description string, t *testing.T, actual error, expected string) {
	SoftAssertEqualErrorStep(description, assert.New(t), actual, expected)
}

func AssertEqualValuesStep(description string, t *testing.T, actual interface{}, expected interface{}) {
	SoftAssertEqualValuesStep(description, assert.New(t), actual, expected)
}

func AssertFalseStep(description string, t *testing.T, actual bool) {
	SoftAssertFalseStep(description, assert.New(t), actual)
}

func AssertTrueStep(description string, t *testing.T, actual bool) {
	SoftAssertTrueStep(description, assert.New(t), actual)
}
