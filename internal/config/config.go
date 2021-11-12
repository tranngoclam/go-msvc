package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/sethvargo/go-envconfig"
)

var (
	ErrProcessConfig = errors.New("config: cannot load config from env")
)

// Config defines configuration using composite pattern
type Config struct {
	sv  Server
	db  Database
	tel Telemetry
}

type Server struct {
	Port int `env:"SERVER_PORT"`
}

type Database struct {
	Host     string `env:"DB_HOST, default=localhost"`
	Port     int    `env:"DB_PORT, default=3306"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD" json:"-"`
	Name     string `env:"DB_NAME"`
}

type Telemetry struct {
	PprofPort int `env:"PPROF_PORT, default=-1"`
}

// New loads configuration from env and return Config,
// default values are used unless the env are set
func New(ctx context.Context) (Config, error) {
	var (
		cfg Config
		db  Database
		tel Telemetry
	)

	if err := envconfig.Process(ctx, &db); err != nil {
		return cfg, fmt.Errorf("%+v, telemetry", ErrProcessConfig)
	}
	if err := envconfig.Process(ctx, &tel); err != nil {
		return cfg, fmt.Errorf("%+v, telemetry", ErrProcessConfig)
	}

	cfg.db = db
	cfg.tel = tel

	return cfg, nil
}

func (c Config) Database() Database {
	return c.db
}

func (c Config) Telemetry() Telemetry {
	return c.tel
}
