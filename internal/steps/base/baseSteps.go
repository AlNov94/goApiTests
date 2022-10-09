package baseSteps

import (
	"fmt"

	"github.com/dailymotion/allure-go"
)

//allure.Step wrapper for functions with returning types and custom step parameters

func Step[T any](description string, action func() T, stepParameters ...interface{}) T {
	var result T
	allure.StepWithParameter(description, getStepParameters(stepParameters...), func() {
		result = action()
	})
	return result
}

//allure.Step wrapper for functions without returning types and custom step parameters

func StepNoResult(description string, action func(), stepParameters ...interface{}) {
	allure.StepWithParameter(description, getStepParameters(stepParameters...), func() {
		action()
	})
}

func getStepParameters(values ...interface{}) map[string]interface{} {
	parameters := make(map[string]interface{})
	len := len(values)
	for i := 0; i < len; i = i + 2 {
		s := fmt.Sprintf("%v", values[i])
		parameters[s] = values[i+1]
	}
	return parameters
}
