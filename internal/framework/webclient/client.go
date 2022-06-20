package webclient

import (
	"fmt"
	"goApiTests/goApiTests/internal/framework/webclient/decoders"
	"goApiTests/goApiTests/internal/framework/webclient/encoders"
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
	resp, err := httpClient.Do(req)
	reqBodyOut, _ := client.encoder.ToOutput(request.body, "", "  ")
	log15.Debug(fmt.Sprintf("Executed %s %s body = %s cookies = %s headers = %s", httpMethod, req.URL, reqBodyOut, req.Cookies(), req.Header))
	if resp != nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		client.decoder.Decode(body, &responseBody)
		respHeaders := client.requestFactory.headersToMap(resp)
		respCookies := client.requestFactory.cookieToMap(resp)
		log15.Debug(fmt.Sprintf("Response status code %d body = %s headers = %v cookies = %v", resp.StatusCode, string(body), respHeaders, respCookies))
		return Response{StatusCode: resp.StatusCode, headers: respHeaders, cookies: respCookies}
	}
	log15.Debug(err.Error())
	return Response{StatusCode: 404, Error: err}
}
