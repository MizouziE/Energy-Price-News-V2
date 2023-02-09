package main

import (
	// "fmt"
	"fmt"
	"log"
	"strconv"
	"strings"

	// "strconv"

	// "github.com/MizouziE/database"
	// "github.com/MizouziE/routes"
	"github.com/MizouziE/models"
	"github.com/gocolly/colly"
	// "github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	articles := make([]models.Article, 0)

	c := colly.NewCollector(
		colly.AllowedDomains("news.sky.com"),
	)

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")

		article := models.Article{
			Url:   link,
			Title: strings.TrimSpace(h.Text),
		}
		if strings.Contains(article.Title, "energy") || strings.Contains(article.Title, "climate") || strings.Contains(article.Title, "Energy") || strings.Contains(article.Title, "Climate") {
			articles = append(articles, article)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL: ", r.Request.URL, " failed with response: ", r, "\nError: ", err)
	})

	c.Visit("https://news.sky.com")

	results := unique(articles)

	log.Println("Articles found: " + strconv.Itoa(len(results)) + "\n\n")

	for _, a := range results {
		if strings.Contains(a.Url, "http") {
			// fmt.Printf("Link: %s\nText: %s\n", a.Url, a.Title)
			fmt.Printf("%+v\n\n", a)
		} else {
			// fmt.Printf("Link: https://news.sky.com%s\nText: %s\n", a.Url, a.Title)
			fmt.Printf("%+v\n\n", a)
		}
	}
	// database.Connect()

	// app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	// routes.Setup(app)

	// app.Listen(":8000")
}

func unique(sample []models.Article) []models.Article {
	var unique []models.Article
	type key struct{ value1, value2 string }
	m := make(map[key]int)
	for _, v := range sample {
		k := key{v.Url, v.Title}
		if i, ok := m[k]; ok {
			// Overwrite previous value per requirement in
			// question to keep last matching value.
			unique[i] = v
		} else {
			// Unique key found. Record position and collect
			// in result.
			m[k] = len(unique)
			unique = append(unique, v)
		}
	}
	return unique
}
