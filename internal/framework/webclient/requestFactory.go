package webclient

import (
	"bytes"
	"net/http"
	"strings"
)

var requestFactoryInstance requestFactory

type requestFactory struct{}

func getrequestFactoryInstance() requestFactory {
	return requestFactoryInstance
}

func (requestFactory requestFactory) createRequest(r *Request, httpMethod string, client *webClient) *http.Request {
	postBody, err := client.encoder.Encode(r.body)
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

func (requestFactory requestFactory) fillUrl(r *Request) string {
	var sb strings.Builder
	sb.WriteString(r.baseUrl + r.path)
	l := len(r.queryParams)
	if l != 0 {
		sb.WriteString("?")
	}
	for k, v := range r.queryParams {
		sb.WriteString(k + "=" + v + "&")
	}
	return sb.String()
}

func (requestFactory requestFactory) cookieToMap(resp *http.Response) map[string]string {
	cookieMap := make(map[string]string)
	cookies := resp.Cookies()
	for i := range cookies {
		cookieMap[cookies[i].Name] = cookies[i].Value
	}
	return cookieMap
}

func (requestFactory requestFactory) headersToMap(resp *http.Response) map[string][]string {
	headerMap := make(map[string][]string)
	for name, value := range resp.Header {
		headerMap[name] = value
	}
	return headerMap
}
