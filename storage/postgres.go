package storage

import (
	"fmt"
	"imazine/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct{
	Host		string
	Port		string
	Password	string
	User		string
	DBName		string
	SSLMode		string
}

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(models.ArticleCategory{})
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Article{})
}

func ConnectDB(config *Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database. \n" + err.Error())
	}

	// MigrateModels(db)

	DB = Dbinstance{
		Db: db,
	}
}
