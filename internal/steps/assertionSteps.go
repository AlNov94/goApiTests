package steps

import (
	"goApiTests/internal/dto"
	"goApiTests/internal/entity"
	baseSteps "goApiTests/internal/steps/base"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AssertionSteps struct{}

// hard assert example
func (assertionSteps AssertionSteps) CheckUserFirstName(t *testing.T, user entity.User, expectedName string) {
	baseSteps.AssertEqualStep("Check user first name", t, user.FirstName, expectedName)
}

// soft assert example
func (assertionSteps AssertionSteps) CheckUserResponseFirstName(t *testing.T, user dto.UserResponse, expectedName string) {
	baseSteps.StepNoResult("Check user response",
		func() {
			assert.Equal(t, user.FirstName, expectedName)
			baseSteps.AssertAll(t)
		},
		baseSteps.Parameter{Name: "expectedName", Value: expectedName},
		baseSteps.Parameter{Name: "actualsUser", Value: user})
}
