package main

import (
	"github.com/gofiber/fiber/v2"
)

type StatusOk struct {
	Status string `json:"status"`
}

type Information []struct {
	ID         string   `json:"_id"`
	Index      int      `json:"index"`
	GUID       string   `json:"guid"`
	IsActive   bool     `json:"isActive"`
	Balance    string   `json:"balance"`
	Picture    string   `json:"picture"`
	Age        int      `json:"age"`
	EyeColor   string   `json:"eyeColor"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	Tags       []string `json:"tags"`
	Friends    []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"friends"`
	Greeting      string `json:"greeting"`
	FavoriteFruit string `json:"favoriteFruit"`
}

var (
	handlerGet = func(c *fiber.Ctx) error {
		c.Set("Connection", "Keep-Alive")
		c.Set("Keep-Alive", "timeout=72")
		return c.JSON(StatusOk{"OK"})
	}

	handlerPost = func(c *fiber.Ctx) error {
		data := &Information{}

		if err := c.BodyParser(data); err != nil {
			return err
		}
		c.Set("Connection", "Keep-Alive")
		c.Set("Keep-Alive", "timeout=72")
		return c.JSON(data)
	}
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:                   false,
		CaseSensitive:             true,
		StrictRouting:             true,
		DisableDefaultDate:        true,
		DisableStartupMessage:     true,
		DisableHeaderNormalizing:  true,
		DisableDefaultContentType: true,
	})

	app.Get("/", handlerGet)
	app.Post("/", handlerPost)

	app.Listen(":3000")
}
