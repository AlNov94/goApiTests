package webclient

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

func (r Request) WithBaseUrl(baseUrl string) Request {
	r.baseUrl = baseUrl
	return r
}

func (r Request) WithPath(path string) Request {
	r.path = path
	return r
}

func (r Request) WithHeaders(headers map[string]string) Request {
	if len(r.headers) == 0 {
		r.headers = headers
		return r
	}
	for k, v := range headers {
		r.headers[k] = v
	}
	return r
}

func (r Request) WithQueryParams(queryParams map[string]string) Request {
	r.queryParams = queryParams
	return r
}

func (r Request) WithBody(body interface{}) Request {
	r.body = body
	return r
}

func (r Request) WithCookies(cookies map[string]string) Request {
	r.cookies = cookies
	return r
}

func (r Request) AddCookies(cookies map[string]string) Request {
	if len(r.cookies) == 0 {
		r.cookies = cookies
		return r
	}
	for k, v := range cookies {
		r.cookies[k] = v
	}
	return r
}
