package property

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/magiconair/properties"
	"os"
	"strings"
)

var PropertyManagerInstance PropertyManager

type PropertyManager struct {
	properties properties.Properties
	cliArgs    map[string]string
}

func init() {
	property := properties.MustLoadFile(os.Getenv("GOPATH")+"/src/restApiTests/application.properties", properties.UTF8)
	PropertyManagerInstance = PropertyManager{properties: *property, cliArgs: getArgs()}
}

func GetPropertyManagerInstance() PropertyManager {
	return PropertyManagerInstance
}

func (propertyManager PropertyManager) GetProperty(propertyName string) string {
	property, ok := propertyManager.cliArgs[propertyName]
	if ok == false {
		return property
	}
	properties, ok := propertyManager.properties.Get(propertyName)
	if ok == false {
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
