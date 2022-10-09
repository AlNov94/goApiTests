package webclient

const (
	GET     HttpMethod = "GET"
	POST    HttpMethod = "POST"
	PUT     HttpMethod = "PUT"
	PATCH   HttpMethod = "PATCH"
	DELETE  HttpMethod = "DELETE"
	TRACE   HttpMethod = "TRACE"
	OPTIONS HttpMethod = "OPTIONS"
	HEAD    HttpMethod = "HEAD"
	CONNECT HttpMethod = "CONNECT"
)

type HttpMethod string
