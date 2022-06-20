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

func (ms mockService) GetMockUser(body *dto.UserResponse) webclient.Response {
	queryParams := make(map[string]string)
	queryParams["suffix"] = "1"
	r := webclient.NewRequest().WithBaseUrl("http://localhost:8090").WithPath("/mock-user").WithQueryParams(queryParams)
	return webclient.NewClient(decoders.JsonDecoder{}, encoders.JsonEncoder{}).DoRequest(webClientEnums.GET, &r, body)
}
