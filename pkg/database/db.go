package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"intern_L0/pkg/order"
	"os"
)

var Db *gorm.DB

const Dsn = "user=student password=12345 dbname=l0 port=5432 sslmode=disable TimeZone=Europe/Moscow"

func InitDb(dsn string) {
	var err error
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	var dsnh = "host=" + dbHost + " " + dsn
	Db, err = gorm.Open(postgres.Open(dsnh), &gorm.Config{})

	if err != nil {
		panic("Подключение к БД не установлено")
	}
	Migrate()
}

func Migrate() {
	Db.AutoMigrate(&order.Order{})
	Db.AutoMigrate(&order.Delivery{})
	Db.AutoMigrate(&order.Payment{})
	Db.AutoMigrate(&order.Items{})
}
