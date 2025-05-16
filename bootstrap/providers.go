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

func ProvideDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
		return nil, err
	}
	if os.Getenv("DATABASE_DEBUG") == "true" {
		db = db.Debug()
	}
	return db, nil
}

var BaseSet = wire.NewSet(ProvideDB, ProvideLogger)
