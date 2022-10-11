package steps

import (
	"goApiTests/internal/api"
	"goApiTests/internal/dto"
	"goApiTests/internal/framework/webclient"
	baseSteps "goApiTests/internal/steps/base"
)

//api steps example

type MockServiceApiSteps struct{}

func (mockserviceApiSteps MockServiceApiSteps) GetMockUserStep(suffix string, responseBody *dto.UserResponse) webclient.Response {
	return baseSteps.Step("GetMockUser call",
		func() webclient.Response {
			return api.GetMockServiceApi().GetMockUser(suffix, responseBody)
		},
		baseSteps.Parameter{Name: "suffix", Value: suffix})
}
