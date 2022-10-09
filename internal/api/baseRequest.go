package api

import (
	"goApiTests/internal/framework/webclient"
	"goApiTests/internal/framework/webclient/decoders"
	"goApiTests/internal/framework/webclient/encoders"
)

func getBaseRequest() webclient.Request {
	return webclient.NewRequest().WithBaseUrl("http://localhost:8090").WithContentType(webclient.JSON)
}

func getBaseClient() webclient.WebClient {
	return webclient.NewClient(decoders.JsonDecoder{}, encoders.JsonEncoder{})
}
