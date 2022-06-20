package main

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"goApiTests/goApiTests/internal/framework/webclient/encoders"
	"goApiTests/goApiTests/internal/repository"
)

func main() {
	fmt.Println(1111)
	fmt.Println(encoders.JsonEncoder{})
	log15.Debug("111")
	u := repository.GetUserRepository().FindUserByID(1)
	fmt.Println(u)
}
