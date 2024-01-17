package main

import (
	"fmt"
	"net/http"

	"github.com/andersonigorf/goexpert-rate-limiter/configs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	fmt.Println("Starting web server on port", configs.WebServerPort)
	http.ListenAndServe(configs.WebServerPort, r)
}
