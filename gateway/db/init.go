package db

import (
	"fmt"
	"gateway/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var db *gorm.DB

func Init() {
	var err error
	c := utils.GetConf().Postgres
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.Host, c.Port, c.User, c.DbName, c.Password)
	db, err = gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	if c.ShowSQL == "yes" {
		db = db.Debug()
	}
	if err = db.AutoMigrate(
		&ClusterInfo{},
	).Error; err != nil {
		panic(err)
	}
}
