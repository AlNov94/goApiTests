package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/magiconair/properties"
)

var ConfigInstance config

type config struct {
	properties properties.Properties
	cliArgs    map[string]string
}

func init() {
	property := properties.MustLoadFile(os.Getenv("GOPATH")+"\\src\\goApiTests\\application.properties", properties.UTF8)
	ConfigInstance = config{properties: *property, cliArgs: getArgs()}
}

func GetConfigInstance() config {
	return ConfigInstance
}

func (pm config) GetProperty(propertyName string) string {
	property, ok := pm.cliArgs[propertyName]
	if ok {
		return property
	}
	properties, ok := pm.properties.Get(propertyName)
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
