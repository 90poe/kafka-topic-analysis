package env

import (
	"fmt"
	"github.com/90poe/service-chassis/logging"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
)

//VERSION of the Program
const VERSION = "v0.1.0"

type Config struct {
	PrettyLogOutput bool
	LogLevel        logging.Level
	OutputConfig    bool
	HealthcheckPort int
	JSONFilePath    string
}

var Settings *Config

func (c *Config) verifyOperations() { //nolint
	var errorMsg string

	if len(c.JSONFilePath) == 0 {
		errorMsg = "The path to the json output from kafkacat, for analysis"
	}

	if len(errorMsg) != 0 {
		log.Println(errorMsg)
		pflag.PrintDefaults()
		os.Exit(1)
	}
}
func init() {

	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")
	viper.SetDefault("PRETTY_LOG_OUTPUT", true)
	viper.SetDefault("LOG_LEVEL", "DEBUG")
	viper.SetDefault("OUTPUT_CONFIG", false)
	viper.SetDefault("HEALTHCHECK_PORT", 8888)

	version := pflag.BoolP("version", "v", false, "prints the current version")
	// Compulsory arg flags
	jsonFilepath := pflag.StringP("json-filepath", "j", "", "the path to the json file")

	logLevel, err := logging.ParseLevel(viper.GetString("LOG_LEVEL"))
	if err != nil {
		panic(fmt.Sprintf("failed to parse log level: %v", err))
	}
	if *version {
		fmt.Printf("Version: %s\n", VERSION)
		os.Exit(0)
	}

	Settings = &Config{
		PrettyLogOutput: viper.GetBool("PRETTY_LOG_OUTPUT"),
		LogLevel:        *logLevel,
		OutputConfig:    viper.GetBool("OUTPUT_CONFIG"),
		HealthcheckPort: viper.GetInt("HEALTHCHECK_PORT"),
		JSONFilePath:    *jsonFilepath,
	}
}
