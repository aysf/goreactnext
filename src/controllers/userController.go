package controllers

import (
	"github.com/aysf/goreactnext/src/database"
	"github.com/aysf/goreactnext/src/models"
	"github.com/gofiber/fiber/v2"
)

func Ambassadors(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_ambasador = true").Find(&users)

	return c.JSON(users)
}

func Ranking(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users, models.User{
		IsAmbasador: true,
	})

	var result []interface{}

	for _, user := range users {
		ambassador := models.Ambassador(user)
		ambassador.CalculateRevenue(database.DB)

		result = append(result, fiber.Map{
			user.Name(): ambassador.Revenue,
		})

	}

	return c.JSON(result)
}
