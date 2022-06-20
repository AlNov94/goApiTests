package main

import (
	"fmt"
	"goApiTests/goApiTests/internal/framework/webclient/encoders"
	"goApiTests/goApiTests/internal/repository"

	"github.com/inconshreveable/log15"
)

func main() {
	fmt.Println(1111)
	fmt.Println(encoders.JsonEncoder{})
	log15.Debug("111")
	u := repository.GetUserRepository().FindUserByID(102)
	fmt.Println(u)
	repository.GetUserRepository().DeleteUser(&u)
}
