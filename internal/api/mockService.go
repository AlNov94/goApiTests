package api

import (
	"goApiTests/goApiTests/internal/dto"
	"goApiTests/goApiTests/internal/framework/webclient"
	"goApiTests/goApiTests/internal/framework/webclient/decoders"
	"goApiTests/goApiTests/internal/framework/webclient/encoders"
	"goApiTests/goApiTests/internal/framework/webclient/webClientEnums"
)

type mockService struct{}

var mockServiceInstance mockService = mockService{}

func GetMockService() mockService {
	return mockServiceInstance
}

func (ms mockService) GetMockUser(suffix string, responseBody *dto.UserResponse) webclient.Response {
	queryParams := make(map[string]string)
	queryParams["suffix"] = suffix
	r := getBaseRequest().WithPath("/mock-user").WithQueryParams(queryParams)
	return webclient.NewClient(decoders.JsonDecoder{}, encoders.JsonEncoder{}).DoRequest(webClientEnums.GET, &r, responseBody)
}
