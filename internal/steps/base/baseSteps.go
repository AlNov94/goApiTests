package baseSteps

import (
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

//allure.Step wrapper for functions with returning types and custom step parameters

func Step[T any](description string, action func() T, stepParameters ...Parameter) T {
	var result T
	allure.StepWithParameter(description, getStepParameters(stepParameters...), func() {
		result = action()
	})
	return result
}

//allure.Step wrapper for functions without returning types and custom step parameters

func StepNoReturn(description string, action func(), stepParameters ...Parameter) {
	allure.StepWithParameter(description, getStepParameters(stepParameters...), func() {
		action()
	})
}

// assert all soft assertions
func AssertAll(t *testing.T) {
	passed := assert.False(t, t.Failed())
	if passed == false {
		t.FailNow()
	}
}

type Parameter struct {
	Name  string
	Value interface{}
}

func getStepParameters(parameters ...Parameter) map[string]interface{} {
	parametersMap := make(map[string]interface{})
	len := len(parameters)
	for i := 0; i < len; i++ {
		parametersMap[parameters[i].Name] = parameters[i].Value
	}
	return parametersMap
}
