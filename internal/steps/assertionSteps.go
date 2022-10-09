package steps

import (
	"fmt"
	"goApiTests/internal/dto"
	"goApiTests/internal/entity"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

type AssertionSteps struct{}

func (assertionSteps AssertionSteps) CheckUserFirstName(t *testing.T, user entity.User, expectedName string) {
	allure.Step(fmt.Sprintf("Check user first name to be equal %s", expectedName), func() {
		assert := assert.New(t)
		assert.Equal(user.FirstName, expectedName)
	})
}

func (assertionSteps AssertionSteps) CheckUserResponseFirstName(t *testing.T, user dto.UserResponse, expectedName string) {
	allure.Step(fmt.Sprintf("Check get user response first name to be equal %s", expectedName), func() {
		assert := assert.New(t)
		assert.Equal(user.FirstName, expectedName)
	})
}
