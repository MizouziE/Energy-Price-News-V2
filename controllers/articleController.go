package controllers

import (
	"github.com/MizouziE/database"
	"github.com/MizouziE/models"
	"github.com/gofiber/fiber/v2"
)

func Articles(c *fiber.Ctx) error {
	db := database.DB
	var articles []models.Article
	db.Preload("Source").Find(&articles)

	return c.JSON(articles)
}
