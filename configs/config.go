package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	Mode        string    `yaml:"mode"`
	ServiceName string    `yaml:"serviceName"`
	Env         string    `yaml:"env"`
	APIs        Apis      `yaml:"apis"`
	GrpcAPIs    GrpcApis  `yaml:"grpcApis"`
	DBConfigs   DBConfigs `yaml:"dbConfigs"`
	DBNames     DBNames   `yaml:"dbNames"`
	DBPassword  string    `yaml:"dbPassword"`
	Migration   bool      `yaml:"enableMigration"`
}

type DBConfigs struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type DBNames struct {
	V1 string `yaml:"v1"`
}

type Apis struct {
	Address         string        `yaml:"address"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout"`
}

type GrpcApis struct {
	Address         string        `yaml:"address"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout"`
}

type Logging struct {
	IsDBLogFormatted bool `yaml:"isDBLogFormatted"`
}

func GetConfig() *Config {
	conf := Config{}

	mode := os.Getenv(Mode)
	if mode == "" {
		mode = ModeLocal
	}
	conf.Mode = mode
	file := fmt.Sprintf(ConfigFileFormat, mode)

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigName(file)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workingDir + "/../../")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal(err)
	}

	return &conf
}
