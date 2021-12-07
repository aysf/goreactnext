package database

import (
	"github.com/aysf/goreactnext/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambasador"), &gorm.Config{})

	if err != nil {
		panic("could not connect with the database")
	}

}

func Automigrate() {
	DB.AutoMigrate(models.User{})
}
