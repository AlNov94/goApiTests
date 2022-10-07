package tests

import (
	"os"
	"testing"

	"github.com/inconshreveable/log15"
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
		log15.Debug(err.Error())
		os.Exit(1)
	}
	return path
}
