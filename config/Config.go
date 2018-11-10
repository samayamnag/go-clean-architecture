package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var environments = map[string]string{
	"development": "config/config.json",
	"production":  "config/config.json",
	"testing":     "config/test.json",
}

type Config struct {
	AppName         string
	AppDebug        bool
	APIVersion      string
	ServerPort      string
	MagicalDate     string
	MongoDBHost     string
	MongoDBPort     int
	MongoDBName     string
	MongoDBUsername string
	MongoDBPassword string
	MongoDBConnPool int
}

var config Config = Config{}
var env = "production"
var debug = true

func Init() {
	env = os.Getenv("GO_ICMYC_ENV")
	envDebug := os.Getenv("GO_ICMYC_DEBUG")
	if env == "" {
		fmt.Println("Warning: Setting production environment due to lack of GO_ICMYC_ENV value")
		env = "production"
	}
	if envDebug == "" {
		fmt.Println("Warning: Setting production environment to DEGUB MODE due to lack of GO_ICMYC_DEBUG value")
		debug = true
	} else {
		envDebug, err := strconv.ParseBool(envDebug)
		if err != nil {
			fmt.Println("Warning: Not able to parese GO_ICMYC_DEBUG value")
		} else {
			debug = envDebug
		}
	}
	LoadConfigByEnv(env)
}

func LoadConfigByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	config = Config{}
	jsonErr := json.Unmarshal(content, &config)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

func Get() Config {
	if &config == nil {
		Init()
	}
	return config
}

func GetEnvironment() string {
	return env
}

func GetAPIVersion() string {
	return config.APIVersion
}

func DegugModeEnabled() bool {
	return debug
}

func IsTestEnvironment() bool {
	return env == "testing"
}
