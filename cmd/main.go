package main

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"goApiTests/goApiTests/internal/framework/webclient/encoders"
)

func main() {
	fmt.Println(1111)
	log15.Debug("111", encoders.JsonEncoder{})
}
