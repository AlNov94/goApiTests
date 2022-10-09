package tests

import (
	"os"
	"sync"

	"github.com/inconshreveable/log15"
)

func init() {
	var once sync.Once
	once.Do(func() {
		prepareAllure()
	})
}

func prepareAllure() {
	log15.Debug("<---------------------- Prepare allure")
	directory := getDirectory()
	allureDirectory := directory + "/allure-results"
	os.RemoveAll(allureDirectory)
	os.Mkdir(allureDirectory, os.ModePerm)
	os.Setenv("ALLURE_RESULTS_PATH", directory)
	log15.Debug("<---------------------- Allure prepared")
}

func getDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		log15.Debug(err.Error())
		os.Exit(1)
	}
	return path
}
