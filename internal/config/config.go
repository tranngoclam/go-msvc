package config

import (
	"context"
	"errors"
	"github.com/sethvargo/go-envconfig"
)

var (
	ErrProcessConfig = errors.New("config: cannot load config from env")
)

type Config struct {
	db Database
}

type Database struct {
	Host     string `env:"DB_HOST, default=localhost"`
	Port     int    `env:"DB_PORT, default=3306"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD" json:"-"`
	Name     string `env:"DB_NAME"`
}

// New loads configuration from env and return Config,
// default values are used unless the env are set
func New(ctx context.Context) (Config, error) {
	var (
		cfg Config
		db  Database
	)
	if err := envconfig.Process(ctx, &db); err != nil {
		return cfg, ErrProcessConfig
	}

	cfg.db = db

	return cfg, nil
}

func (c Config) Database() Database {
	return c.db
}
