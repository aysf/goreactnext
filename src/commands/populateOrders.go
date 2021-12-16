package main

import (
	"math/rand"

	"github.com/aysf/goreactnext/src/database"
	"github.com/aysf/goreactnext/src/models"
	faker "github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {

		var orderItems []models.OrderItem

		for j := 0; j < rand.Intn(5); j++ {
			price := float64(rand.Intn(90) + 10)
			qty := uint(rand.Intn(5))

			orderItems = append(orderItems, models.OrderItem{
				ProductTitle:     faker.Word(),
				Price:            price,
				Quantity:         qty,
				AdminRevenue:     0.9 * price * float64(qty),
				AmbasadorRevenue: 0.9 * price * float64(qty),
			})
		}

		database.DB.Create(&models.Order{
			UserId:          uint(rand.Intn(30) + 1),
			Code:            faker.Username(),
			AmbassadorEmail: faker.Email(),
			FirstName:       faker.FirstName(),
			LastName:        faker.LastName(),
			Complete:        true,
			OrderItems:      orderItems,
		})

	}

}
