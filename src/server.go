package main

import (
	"github.com/gofiber/fiber"
	"github.com/lukasvdberk/go-url-shortner/database"
	"strconv"
	"strings"
)

func main() {
	app := fiber.New()

	app.Post("/save_url", func(c *fiber.Ctx) {
		//
		shortUrl := new(database.ShortUrl)
		shortUrl.RealUrl = c.FormValue("url_to_shorten")

		shortUrl = database.SaveUrl(shortUrl)

		if err := c.JSON(shortUrl); err != nil {
			// Sends error if it fails
			c.Status(500).Send(err)
			return
		}
	})

	app.Get("/redirect/:db_id", func(c *fiber.Ctx) {
		id, err := strconv.ParseInt(
			strings.TrimSpace(c.Params("db_id")), 10, 64)

		if err == nil {
			shortUrl := database.GetShortUrlById(id)
			if shortUrl != nil {
				c.Redirect(shortUrl.RealUrl)
			} else {
				c.Status(500).Send("Id not found")
			}
		} else {
			// Sends error if it fails
			c.Status(500).Send("Id is not a valid integer")
		}
	})

	app.Static("/", "./public")

	_ = app.Listen("0.0.0.0:4000")
}
