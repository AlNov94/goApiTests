package api

import (
	"goApiTests/internal/dto"
	"goApiTests/internal/framework/webclient"
	"goApiTests/internal/framework/webclient/decoders"
	"goApiTests/internal/framework/webclient/encoders"
	"goApiTests/internal/framework/webclient/webClientEnums"
)

type mockService struct{}

var mockServiceInstance mockService = mockService{}

func GetMockService() mockService {
	return mockServiceInstance
}

func (ms mockService) GetMockUser(suffix string, responseBody *dto.UserResponse) webclient.Response {
	queryParams := make(map[string]string)
	queryParams["suffix"] = suffix
	request := getBaseRequest().WithPath("/mock-user").WithQueryParams(queryParams)
	return webclient.NewClient(decoders.JsonDecoder{}, encoders.JsonEncoder{}).DoRequest(webClientEnums.GET, &request, responseBody)
}
