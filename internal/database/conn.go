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
func New(ctx context.Context, dbCfg config.Database) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	gDB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, ErrOpenDB
	}

	db := &DB{db: gDB}
	err = db.Ping(ctx)
	if err != nil {
		return nil, ErrPingDB
	}

	return db, nil
}

// Ping performs a ping command to the database
func (db *DB) Ping(ctx context.Context) error {
	sDB, err := db.db.DB()
	if err != nil {
		return err
	}

	return sDB.PingContext(ctx)
}

// Close retrieves sql/database instance and closes all connections to the database
func (db *DB) Close() error {
	sDB, err := db.db.DB()
	if err != nil {
		return err
	}

	return sDB.Close()
}
