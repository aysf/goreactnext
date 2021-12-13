package main

import (
	"github.com/aysf/goreactnext/src/database"
	"github.com/aysf/goreactnext/src/models"
	faker "github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FirstName:   faker.FirstName(),
			LastName:    faker.LastName(),
			Email:       faker.Email(),
			IsAmbasador: true,
		}

		ambassador.SetPassword("1234")

		database.DB.Create(&ambassador)
	}

}
