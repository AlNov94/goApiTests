package main

import (
	"fmt"
	"goApiTests/goApiTests/framework/webclient/encoders"

	"github.com/inconshreveable/log15"
)

func main() {
	fmt.Println(1111)
	log15.Debug("111", encoders.JsonEncoder{})
}
