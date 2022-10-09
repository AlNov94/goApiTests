package steps

import (
	"fmt"
	"goApiTests/internal/api"
	"goApiTests/internal/dto"
	"goApiTests/internal/framework/webclient"

	"github.com/dailymotion/allure-go"
)

type MockServiceApiSteps struct{}

func (mockserviceApiSteps MockServiceApiSteps) GetMockUserStep(suffix string, responseBody *dto.UserResponse) webclient.Response {
	var result webclient.Response
	allure.Step(fmt.Sprintf("GetMockUser call with suffix %s", suffix), func() {
		result = api.GetMockServiceApi().GetMockUser(suffix, responseBody)
	})
	return result
}
