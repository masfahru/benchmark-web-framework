package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sort"
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
	handlerGet = func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderConnection, "Keep-Alive")
		c.Response().Header().Set("Keep-Alive", "timeout=72")
		return c.JSON(http.StatusOK, StatusOk{"OK"})
	}

	handlerPost = func(c echo.Context) error {
		data := &Information{}

		if err := c.Bind(data); err != nil {
			return err
		}
		// sort data by age
		sort.Slice(*data, func(i, j int) bool {
			return (*data)[i].Age < (*data)[j].Age
		})
		c.Response().Header().Set(echo.HeaderConnection, "Keep-Alive")
		c.Response().Header().Set("Keep-Alive", "timeout=72")
		return c.JSON(http.StatusOK, data)
	}
)

func main() {
	app := echo.New()
	app.HideBanner = true

	app.GET("/", handlerGet)
	app.POST("/", handlerPost)

	app.Start(":3000")
}
