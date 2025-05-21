package bootstrap

import (
	"fmt"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ProvideLogger(name string) *log.Logger {
	return log.New(os.Stdout, name, log.LstdFlags|log.Lshortfile)
}

func ProvideDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
		return nil, err
	}
	if cfg.Debug {
		db = db.Debug()
	}
	return db, nil
}

var BaseSet = wire.NewSet(ProvideDB, ProvideLogger)
