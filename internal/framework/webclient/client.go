package webclient

import (
	"fmt"
	"goApiTests/internal/framework/webclient/decoders"
	"goApiTests/internal/framework/webclient/encoders"
	"io"
	"net/http"

	"github.com/inconshreveable/log15"
)

type webClient struct {
	decoder        decoders.Decoder
	encoder        encoders.Encoder
	requestFactory requestFactory
}

func NewClient(decoder decoders.Decoder, encoder encoders.Encoder) webClient {
	return webClient{decoder: decoder, encoder: encoder, requestFactory: getrequestFactoryInstance()}
}

func (client webClient) DoRequest(httpMethod string, request *Request, responseBody any) Response {
	req := client.requestFactory.createRequest(request, httpMethod, &client)
	httpClient := &http.Client{}
	response, err := httpClient.Do(req)
	requestBodyOutput, _ := client.encoder.ToOutput(request.body, "", "  ")
	log15.Debug(fmt.Sprintf("Executed %s %s body = %s cookies = %s headers = %s", httpMethod, req.URL, requestBodyOutput, req.Cookies(), req.Header))
	if response != nil {
		defer response.Body.Close()
		body, _ := io.ReadAll(response.Body)
		client.decoder.Decode(body, &responseBody)
		responseHeaders := client.requestFactory.headersToMap(response)
		responseCookies := client.requestFactory.cookieToMap(response)
		log15.Debug(fmt.Sprintf("Response status code %d body = %s headers = %v cookies = %v", response.StatusCode, string(body), responseHeaders, responseCookies))
		return Response{StatusCode: response.StatusCode, headers: responseHeaders, cookies: responseCookies}
	}
	log15.Debug(err.Error())
	return Response{StatusCode: 404, Error: err}
}
