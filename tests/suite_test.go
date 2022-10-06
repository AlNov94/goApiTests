package tests

import (
	"goApiTests/internal/framework/config"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ApiTestSuite struct {
	suite.Suite
}

func TestApiTestSuite(t *testing.T) {
	os.Setenv("ALLURE_RESULTS_PATH", config.GetConfig().GetProperty("ALLURE_RESULTS_PATH"))
	suite.Run(t, new(ApiTestSuite))
}
