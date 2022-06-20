package main

import (
	"fmt"
	"goApiTests/goApiTests/internal/api"
	"goApiTests/goApiTests/internal/dto"
	"goApiTests/goApiTests/internal/framework/webclient/encoders"
	"goApiTests/goApiTests/internal/repository"

	"github.com/inconshreveable/log15"
)

func main() {
	fmt.Println(1111)
	fmt.Println(encoders.JsonEncoder{})
	log15.Debug("111")
	u := repository.GetUserRepository().FindUserByID(1)
	fmt.Println(u)
	var u2 dto.UserResponse
	api.GetMockService().GetMockUser(&u2)
	fmt.Println(u2)
}
