package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ApiTestSuite struct {
	suite.Suite
}

func TestApiTestSuite(t *testing.T) {
	directory := getDirectory()
	os.RemoveAll(directory + "/allure-results")
	os.Setenv("ALLURE_RESULTS_PATH", directory)
	suite.Run(t, new(ApiTestSuite))
}

func getDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return path
}
