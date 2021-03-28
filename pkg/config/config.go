package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

//App struct
type Configuration struct {
	App App 
}

type App struct {
	Maps    Maps    `mapstructure:"maps"`
	Weather Weather `mapstructure:"weather"`
}

type Maps struct {
	URL string `mapstructure:"url"`
	Key string `mapstructure:"key"`
}

type Weather struct {
	URL   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
}

//ReadConfig inicia as configurações
func ReadConfig(configFile string) (*Configuration, error) {

	var app *Configuration

	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("/tmp/")
	viper.AddConfigPath("$HOME/")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&app)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshal config file, %s", err)
	}
	return app, nil

}
