package controllers

import (
	"github.com/MizouziE/database"
	"github.com/MizouziE/models"
	"github.com/gofiber/fiber/v2"
)

func Regions(c *fiber.Ctx) error {
	db := database.DB
	var regions []models.Region
	db.Find(&regions)

	return c.JSON(regions)
}

func SingleRegion(c *fiber.Ctx) error {
	db := database.DB
	regionId := c.Params("regionId")
	var articles []models.Article
	db.Joins("Source").Where("region_id = ?", regionId).Find(&articles)

	return c.JSON(articles)
}

func SingleRegionSources(c *fiber.Ctx) error {
	db := database.DB
	regionId := c.Params("regionId")
	var sources []models.Source
	db.Where("region_id = ?", regionId).Find(&sources)

	return c.JSON(sources)
}
