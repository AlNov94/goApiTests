package webclient

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"goApiTests/goApiTests/framework/webclient/decoders"
	"goApiTests/goApiTests/framework/webclient/encoders"
	"io"
	"net/http"
)

type Client struct {
	decoder        decoders.Decoder
	encoder        encoders.Encoder
	requestFactory requestFactory
}

func NewClient(decoder *decoders.Decoder, encoder *encoders.Encoder) Client {
	return Client{decoder: *decoder, encoder: *encoder, requestFactory: getrequestFactoryInstance()}
}

func (client Client) doRequest(r *Request, body *any, httpMethod string) Response {
	req := client.requestFactory.createRequest(r, httpMethod, &client)
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	reqBodyOut, _ := client.encoder.ToOutput(r.body, "", "  ")
	log15.Debug(fmt.Sprintf("Executed %s %s body = %s cookies = %s headers = %s", httpMethod, req.URL, reqBodyOut, req.Cookies(), req.Header))
	if resp != nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		client.decoder.Decode(body, &body)
		respHeaders := client.requestFactory.headersToMap(resp)
		respCookies := client.requestFactory.cookieToMap(resp)
		log15.Debug(fmt.Sprintf("Response status code %d body = %s headers = %v cookies = %v", resp.StatusCode, string(body), respHeaders, respCookies))
		return Response{StatusCode: resp.StatusCode, headers: respHeaders, cookies: respCookies}
	}
	log15.Debug(err.Error())
	return Response{StatusCode: 404, Error: err}
}
