package postgres

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(os.Getenv("POSTGRES_DSN")))
}
