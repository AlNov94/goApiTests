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

func (response Response) Header(headerName string) string {
	return response.HeaderAt(headerName, 0)
}

func (response Response) HeaderAt(headerName string, index int) string {
	headers, ok := response.headers[headerName]
	if ok == false {
		errMsg := fmt.Sprintf("Header %s not exist", headerName)
		log15.Debug(errMsg)
		panic(errMsg)
	}
	if len(headers) < index+1 {
		errMsg := fmt.Sprintf("Header %s with index %d not exist", headerName, index)
		log15.Debug(errMsg)
		panic(errMsg)
	}
	return headers[index]
}

func (response Response) Cookie(cookieName string) string {
	if len(response.cookies[cookieName]) == 0 {
		errMsg := fmt.Sprintf("Cookie %s not exist", cookieName)
		log15.Debug(errMsg)
		panic(errMsg)
	}
	return response.cookies[cookieName]
}
