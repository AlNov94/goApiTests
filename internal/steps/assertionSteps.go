package steps

import (
	"goApiTests/internal/dto"
	"goApiTests/internal/entity"
	"goApiTests/internal/steps/base"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AssertionSteps struct{}

func (assertionSteps AssertionSteps) CheckUserFirstName(t *testing.T, user entity.User, expectedName string) {
	assert := assert.New(t)
	baseSteps.SoftAssertEqualStep("Check user first name", assert, user.FirstName, expectedName)
	baseSteps.AssertAll(t)
}

func (assertionSteps AssertionSteps) CheckUserResponseFirstName(t *testing.T, user dto.UserResponse, expectedName string) {
	assert := assert.New(t)
	baseSteps.SoftAssertEqualStep("Check get user response first name", assert, user.FirstName, expectedName)
	baseSteps.AssertAll(t)
}
