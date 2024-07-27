package config

import (
	"fmt"
	"os"

	"github.com/DenisCom3/m-chat-server/internal/config/auth"
	"github.com/DenisCom3/m-chat-server/internal/config/grpc"
	"github.com/DenisCom3/m-chat-server/internal/config/postgres"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

var cfg *config

type YamlConfig struct {
	Postgres postgres.Postgres `yaml:"postgres" env-required:"true"`
	Grpc     grpc.Grpc         `yaml:"grpc" env-required:"true"`
	Auth     auth.Auth        `yaml:"auth" env-required:"true"`
}

type config struct {
	postgres Postgres
	grpc     Grpc
	auth     Auth
}

type Postgres interface {
	Dsn() string
}

type Grpc interface {
	Address() string
}

type Auth interface {
	Address() string
}

func GetPostgres() Postgres {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg.postgres
}

func GetGrpc() Grpc {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg.grpc
}

func GetAuth() Auth {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg.auth
}
func MustLoad() error {

	if cfg != nil {
		return fmt.Errorf("config already initialized")
	}

	err := godotenv.Load()

	if err != nil {
		return fmt.Errorf("%w", err)
	}
	configPath := os.Getenv("APP_CONFIG_PATH")

	if configPath == "" {
		return fmt.Errorf("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist: %s", configPath)
	}

	var yaml YamlConfig

	if err := cleanenv.ReadConfig(configPath, &yaml); err != nil {
		return fmt.Errorf("cannot read config: %s", err)
	}

	cfg = &config{
		postgres: yaml.Postgres,
		grpc:     yaml.Grpc,
		auth:     yaml.Auth,
	}
	return nil
}
