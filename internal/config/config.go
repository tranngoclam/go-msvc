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
	srv Server
	db  Database
	tel Telemetry
}

type Server struct {
	Port string `env:"SERVER_PORT"`
}

func (s Server) Addr() string {
	return ":" + s.Port
}

// Database defines configuration for database
type Database struct {
	Host     string `env:"DB_HOST, default=localhost"`
	Port     string `env:"DB_PORT, default=3306"`
	User     string `env:"DB_USER, default=todo"`
	Password string `env:"DB_PASSWORD, default=password" json:"-"`
	Name     string `env:"DB_NAME, default=todo"`
}

// Telemetry defines configuration for telemetry
type Telemetry struct {
	PprofPort string `env:"PPROF_PORT, default=-1"`
}

// New loads configuration from env and return Config,
// default values are used unless the env are set
func New(ctx context.Context) (Config, error) {
	var (
		cfg Config
		srv Server
		db  Database
		tel Telemetry
	)

	if err := envconfig.Process(ctx, &srv); err != nil {
		return cfg, fmt.Errorf("%+v, server", ErrProcessConfig)
	}
	if err := envconfig.Process(ctx, &db); err != nil {
		return cfg, fmt.Errorf("%+v, database", ErrProcessConfig)
	}
	if err := envconfig.Process(ctx, &tel); err != nil {
		return cfg, fmt.Errorf("%+v, telemetry", ErrProcessConfig)
	}

	cfg.srv = srv
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

func (c Config) Server() Server {
	return c.srv
}
