package controllers

import (
	"github.com/MizouziE/database"
	"github.com/MizouziE/models"
	"github.com/gofiber/fiber/v2"
)

func Keywords(c *fiber.Ctx) error {
	db := database.DB
	var keywords []models.Keyword
	db.Find(&keywords)

	return c.JSON(keywords)
}
