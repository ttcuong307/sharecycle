package configs

import (
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME" envDefault:"sharecycle"`
	Env         string `env:"ENV" envDefault:"local"`
	API         struct {
		Address         string        `env:"API_ADDRESS" envDefault:"0.0.0.0:8080"`
		ShutdownTimeout time.Duration `env:"API_SHUTDOWN_TIMEOUT" envDefault:"60s"`
	}
	Database struct {
		UserName string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Host     string `env:"DB_HOST"`
		Port     int    `env:"DB_PORT" envDefault:"3306"`
		DBName   string `env:"DB_NAME" envDefault:"sharecycle"`
		Params   string `env:"DB_PARAMS_OVERRIDES" envDefault:"sslmode=disable"`
	}
	// DBConfigs *DBConfigs `yaml:"dbConfigs"`
	// DBNames   *DBNames   `yaml:"dbNames"`
	Migration struct {
		Enable bool `env:"ENABLE_MIGRATE"`
	}
}

type DBConfigs struct {
	Region   string `yaml:"region"`
	IdleConn int    `yaml:"idleConn"`
	MaxConn  int    `yaml:"maxConn"`
	Debug    bool   `yaml:"debug"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type DBNames struct {
	V1 string `yaml:"v1"`
}

func PopulateENV() error {
	err := godotenv.Load("local.env")
	if err != nil {
		return err
	}
	return nil
}

type Logging struct {
	IsDBLogFormatted bool `yaml:"isDBLogFormatted"`
}
