package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"ucenter/utils"
)

var db *gorm.DB

func Init() {
	var err error
	c := utils.GetConf().Mysql
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", c.User, c.Pwd, c.Host, c.Port, c.Db)
	db, err = gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	if err = db.Debug().AutoMigrate(
		&SubSystem{}, &PermissionInfo{}, &RoleInfo{}, &RelationRolePermission{},
		&UserInfo{}, &UserGroup{}, &RelationUserGroup{}, &RelationUserRolePermission{}, &RelationGroupRolePermission{},
	).Error; err != nil {
		panic(err)
	}
}
