package controllers

import (
	"github.com/aysf/goreactnext/src/database"
	"github.com/aysf/goreactnext/src/models"
	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Find(&orders)

	return c.JSON(orders)
}
