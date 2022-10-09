package webclient

//representing http request struct

type Request struct {
	baseUrl     string
	path        string
	headers     map[string]string
	queryParams map[string]string
	body        interface{}
	cookies     map[string]string
}

func NewRequest() Request {
	return Request{}
}

func (request Request) WithBaseUrl(baseUrl string) Request {
	request.baseUrl = baseUrl
	return request
}

func (request Request) WithPath(path string) Request {
	request.path = path
	return request
}

func (request Request) WithContentType(contentType ContentType) Request {
	headers := make(map[string]string)
	headers["Content-Type"] = string(contentType)
	return request.WithHeaders(headers)
}

func (request Request) WithHeaders(headers map[string]string) Request {
	if len(request.headers) == 0 {
		request.headers = headers
		return request
	}
	for k, v := range headers {
		request.headers[k] = v
	}
	return request
}

func (request Request) WithQueryParams(queryParams map[string]string) Request {
	request.queryParams = queryParams
	return request
}

func (request Request) WithBody(body interface{}) Request {
	request.body = body
	return request
}

func (request Request) WithCookies(cookies map[string]string) Request {
	request.cookies = cookies
	return request
}

func (request Request) AddCookies(cookies map[string]string) Request {
	if len(request.cookies) == 0 {
		request.cookies = cookies
		return request
	}
	for k, v := range cookies {
		request.cookies[k] = v
	}
	return request
}
