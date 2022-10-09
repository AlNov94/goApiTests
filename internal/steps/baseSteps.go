package steps

import (
	"fmt"

	"github.com/dailymotion/allure-go"
)

func Step[T any](description string, action func() T, stepParameters ...interface{}) T {
	var result T
	allure.StepWithParameter(description, getStepParameters(stepParameters...), func() {
		result = action()
	})
	return result
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
