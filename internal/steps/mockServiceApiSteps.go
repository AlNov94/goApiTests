package steps

import (
	"goApiTests/internal/api"
	"goApiTests/internal/dto"
	"goApiTests/internal/framework/webclient"
)

type MockServiceApiSteps struct{}

func (mockserviceApiSteps MockServiceApiSteps) GetMockUserStep(suffix string, responseBody *dto.UserResponse) webclient.Response {
	return Step("GetMockUser call",
		func() webclient.Response {
			return api.GetMockServiceApi().GetMockUser(suffix, responseBody)
		},
		"suffix", suffix)
}
