package webclient

import (
	"fmt"
	"goApiTests/internal/framework/webclient/decoders"
	"goApiTests/internal/framework/webclient/encoders"
	"io"
	"net/http"

	"github.com/inconshreveable/log15"
)

type WebClient struct {
	decoder        decoders.Decoder
	encoder        encoders.Encoder
	requestFactory requestFactory
}

func NewClient(decoder decoders.Decoder, encoder encoders.Encoder) WebClient {
	return WebClient{decoder: decoder, encoder: encoder, requestFactory: getRequestFactoryInstance()}
}

func (webClient WebClient) DoRequest(httpMethod string, request *Request, responseBody ...any) Response {
	req := webClient.requestFactory.createRequest(request, httpMethod, &webClient)
	httpClient := &http.Client{}
	response, err := httpClient.Do(req)
	requestBodyOutput, _ := webClient.encoder.ToOutput(request.body, "", "  ")
	log15.Debug(fmt.Sprintf("Executed %s %s body = %s cookies = %s headers = %s", httpMethod, req.URL, requestBodyOutput, req.Cookies(), req.Header))
	if response != nil {
		return webClient.deserializeResponse(response, responseBody...)
	}
	log15.Debug(err.Error())
	return Response{StatusCode: 404, Error: err}
}

func (webClient WebClient) deserializeResponse(response *http.Response, responseBody ...any) Response {
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	responseHeaders := webClient.requestFactory.headersToMap(response)
	responseCookies := webClient.requestFactory.cookieToMap(response)
	if len(responseBody) != 0 {
		webClient.decoder.Decode(body, &responseBody[0])
		log15.Debug(fmt.Sprintf("Response status code %d body = %s headers = %v cookies = %v", response.StatusCode, string(body), responseHeaders, responseCookies))
	} else {
		log15.Debug(fmt.Sprintf("Response status code %d headers = %v cookies = %v", response.StatusCode, responseHeaders, responseCookies))
	}
	return Response{StatusCode: response.StatusCode, headers: responseHeaders, cookies: responseCookies}
}
