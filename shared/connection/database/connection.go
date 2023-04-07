package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Connection(dialect string, DbHost string, DbPort int, DbName string, DbUser string, DbPass string, Profile bool) (*gorm.DB, error) {

	var err error
	var db *gorm.DB

	config := logger.Config{
		Colorful:                  true,
		LogLevel:                  logger.Info,
		SlowThreshold:             time.Second,
		IgnoreRecordNotFoundError: true,
	}

	if !Profile {
		config.LogLevel = logger.Silent
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		config,
	)

	switch dialect {
	// case "mysql":
	// 	db, err = gorm.Open(mysql.New(mysql.Config{DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DbUser, DbPass, DbHost, strconv.Itoa(DbPort), DbName)}), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	case "postgres":
		db, err = gorm.Open(postgres.New(postgres.Config{DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, strconv.Itoa(DbPort), DbUser, DbPass, DbName), PreferSimpleProtocol: true}), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}, Logger: newLogger})
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
