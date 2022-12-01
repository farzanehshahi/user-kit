package config

import (
	"github.com/qiangxue/go-env"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const (
	defaultServerPort = 8080
)

// config represents an application configuration.
type config struct {
	ServerPort int `yaml:"server_port" env:"SERVER_PORT"`

	PostgresDSN string `yaml:"postgres_dsn" env:"POSTGRES_DSN"`
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string) (*config, error) {
	//	default config
	c := config{
		ServerPort: defaultServerPort,
	}

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	// for example: "API_SERVER_PORT", "8080"
	if err = env.New("APP_", log.Printf).Load(&c); err != nil {
		return nil, err
	}

	return &c, err
}
