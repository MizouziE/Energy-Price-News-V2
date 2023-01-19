package controllers

import (
	"strings"

	"github.com/MizouziE/database"
	"github.com/MizouziE/models"
	"github.com/gofiber/fiber/v2"
)

func Sources(c *fiber.Ctx) error {
	db := database.DB
	var sources []models.Source
	db.Find(&sources)

	return c.JSON(sources)
}

func SingleSource(c *fiber.Ctx) error {
	db := database.DB
	sourceId := strings.ToLower(c.Params("sourceId"))
	var articles []models.Article
	db.Where("source_id = ?", sourceId).Find(&articles)

	return c.JSON(articles)
}
