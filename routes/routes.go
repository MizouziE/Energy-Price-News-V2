package routes

import (
	"github.com/MizouziE/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// user endpoints
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	// news endpoints
	app.Get("/api/news", controllers.Articles)
	app.Get("/api/news/sources", controllers.Sources)
	app.Get("/api/news/sources/:sourceId", controllers.SingleSource)
	app.Get("/api/news/regions", controllers.Regions)
	app.Get("/api/news/regions/:regionId", controllers.SingleRegion)
	app.Get("/api/news/regions/:regionId/sources", controllers.SingleRegionSources)
	app.Get("/api/news/keywords", controllers.Keywords)
}
