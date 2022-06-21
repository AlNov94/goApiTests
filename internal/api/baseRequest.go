package api

import "goApiTests/goApiTests/internal/framework/webclient"

func getBaseRequest() webclient.Request {
	return webclient.NewRequest().WithBaseUrl("http://localhost:8090")
}
