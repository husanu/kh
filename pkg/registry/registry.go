package registry

import (
	"fmt"

	t "github.com/audibleblink/kh/pkg/types"
	"github.com/spf13/viper"
)

// Registry holds the Unmarshaled YAML configs where the CLI can dynamically choose which
// service to validate against based on user input.
var Registry = make(map[string]*t.KeyHack)

// Build esatblishes a configuration file search order and parses the YAML file
// into the tool's registry full of KeyHacks
func Build() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("keyhacks")

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/")
	viper.AddConfigPath("$GOPATH/src/github.com/audibleblink/kh/")
	viper.AddConfigPath("/etc/")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.Unmarshal(&Registry)
}
