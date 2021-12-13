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
