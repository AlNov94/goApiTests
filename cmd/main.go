package main

import (
	"fmt"
	"goApiTests/internal/api"
	"goApiTests/internal/dto"
	"goApiTests/internal/framework/webclient/encoders"
	"goApiTests/internal/repository"

	"github.com/inconshreveable/log15"
)

func main() {
	fmt.Println(1111)
	fmt.Println(encoders.JsonEncoder{})
	log15.Debug("111")
	u := repository.GetUserRepository().FindUserById(1)
	fmt.Println(u)
	var u2 dto.UserResponse
	api.GetMockServiceApi().GetMockUser("1", &u2)
	fmt.Println(u2)
}
