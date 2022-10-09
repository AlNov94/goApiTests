package api

import (
	"goApiTests/internal/dto"
	"goApiTests/internal/framework/webclient"
	"goApiTests/internal/framework/webclient/webClientEnums"
)

type mockServiceApi struct{}

var mockServiceApiInstance mockServiceApi = mockServiceApi{}

func GetMockServiceApi() mockServiceApi {
	return mockServiceApiInstance
}

func (mockServiceApi mockServiceApi) GetMockUser(suffix string, responseBody *dto.UserResponse) webclient.Response {
	queryParams := make(map[string]string)
	queryParams["suffix"] = suffix
	request := getBaseRequest().WithPath("/mock-user").WithQueryParams(queryParams)
	return getBaseClient().DoRequest(webClientEnums.GET, &request, responseBody)
}
