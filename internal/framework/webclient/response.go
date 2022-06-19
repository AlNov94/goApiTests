package webclient

import (
	"fmt"
	"github.com/inconshreveable/log15"
)

type Response struct {
	StatusCode int
	Error      error
	headers    map[string][]string
	cookies    map[string]string
}

func (r Response) Header(headerName string) string {
	return r.HeaderAt(headerName, 0)
}

func (r Response) HeaderAt(headerName string, index int) string {
	if len(r.headers[headerName]) == 0 {
		errMsg := fmt.Sprintf("Header %s not exist", headerName)
		log15.Debug(errMsg)
		panic(errMsg)
	}
	if len(r.headers[headerName]) < index+1 {
		errMsg := fmt.Sprintf("Header %s with index %d not exist", headerName, index)
		log15.Debug(errMsg)
		panic(errMsg)
	}
	return r.headers[headerName][index]
}

func (r Response) Cookie(cookieName string) string {
	if len(r.cookies[cookieName]) == 0 {
		errMsg := fmt.Sprintf("Cookie %s not exist", cookieName)
		log15.Debug(errMsg)
		panic(errMsg)
	}
	return r.cookies[cookieName]
}
