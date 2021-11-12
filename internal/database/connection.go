package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/tranngoclam/go-msvc/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrOpenDB = errors.New("db: cannot open database")
	ErrPingDB = errors.New("db: cannot ping to database")
)

type DB struct {
	db *gorm.DB
}

// New constructs gorm db object from database configuration
// TODO: customize gorm configuration
func New(ctx context.Context, cfg config.Config) (*DB, error) {
	dbCfg := cfg.Database()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	gDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, ErrOpenDB
	}

	sDB, err := gDB.DB()
	if err != nil {
		return nil, err
	}

	err = sDB.PingContext(ctx)
	if err != nil {
		return nil, ErrPingDB
	}

	return &DB{db: gDB}, nil
}
