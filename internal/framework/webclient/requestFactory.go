package webclient

import (
	"bytes"
	"net/http"
	"strings"
)

var requestFactoryInstance requestFactory

type requestFactory struct{}

func init() {
	requestFactoryInstance = requestFactory{}
}

func getRequestFactoryInstance() requestFactory {
	return requestFactoryInstance
}

func (requestFactory requestFactory) createRequest(r *Request, httpMethod string, webClient *WebClient) *http.Request {
	postBody, err := webClient.encoder.Encode(r.body)
	if err != nil {
		panic(err)
	}
	requestBody := bytes.NewBuffer(postBody)
	req, err2 := http.NewRequest(httpMethod, requestFactory.fillUrl(r), requestBody)
	if err2 != nil {
		panic(err2)
	}
	for k, v := range r.cookies {
		req.AddCookie(&http.Cookie{Name: k, Value: v})
	}
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}
	return req
}

func (requestFactory requestFactory) fillUrl(request *Request) string {
	var sb strings.Builder
	sb.WriteString(request.baseUrl + request.path)
	l := len(request.queryParams)
	if l != 0 {
		sb.WriteString("?")
	}
	for k, v := range request.queryParams {
		sb.WriteString(k + "=" + v + "&")
	}
	return sb.String()
}

func (requestFactory requestFactory) cookieToMap(response *http.Response) map[string]string {
	cookieMap := make(map[string]string)
	cookies := response.Cookies()
	for i := range cookies {
		cookieMap[cookies[i].Name] = cookies[i].Value
	}
	return cookieMap
}

func (requestFactory requestFactory) headersToMap(response *http.Response) map[string][]string {
	headerMap := make(map[string][]string)
	for name, value := range response.Header {
		headerMap[name] = value
	}
	return headerMap
}
