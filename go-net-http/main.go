package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	handler = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("Keep-Alive", "timeout=72")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(StatusOk{"OK"})
		case "POST":
			data := &Information{}
			reqBody, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(reqBody, data)
			json.NewEncoder(w).Encode(data)
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	}
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
