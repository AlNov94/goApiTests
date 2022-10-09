package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/magiconair/properties"
)

//struct for reading properties

var configInstance config

type config struct {
	properties properties.Properties
	cliArgs    map[string]string
}

func init() {
	property := properties.MustLoadFile(os.Getenv("GOPATH")+"\\src\\goApiTests\\application.properties", properties.UTF8)
	configInstance = config{properties: *property, cliArgs: getArgs()}
}

func GetConfig() config {
	return configInstance
}

func (config config) GetProperty(propertyName string) string {
	property, ok := config.cliArgs[propertyName]
	if ok {
		return property
	}
	properties, ok := config.properties.Get(propertyName)
	if !ok {
		log15.Debug(fmt.Sprintf("Error due to reading property %s", propertyName))
	}
	return properties
}

func getArgs() map[string]string {
	cliArgs := make(map[string]string)
	args := os.Args
	for i := range args {
		arg := args[i]
		if strings.HasPrefix(arg, "-") && strings.Contains(arg, "=") {
			replace := strings.Replace(arg, "-", "", 1)
			before, after, _ := strings.Cut(replace, "=")
			cliArgs[before] = after
		}
	}
	return cliArgs
}
