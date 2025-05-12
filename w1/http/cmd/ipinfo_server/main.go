package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"http/cmd/ipinfo_server/auth"
	"http/cmd/ipinfo_server/db"

	"http/cmd/ipinfo_server/handlers"
	"log"
	"net/http"
)

func main() {

	//db init
	db.InitDB()
	defer db.DB.Close()

	//router init
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	//public path
	router.Post("/register", handlers.RegisterHandler)

	//with token path
	router.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Get("/self_ip", handlers.SelfIpHandler)
		r.Get("/ext_ip/{ip}", handlers.ExtIpHandler)
		r.Get("/history", handlers.HistoryHandler)
	})

	////with basic auth path
	//router.Group(func(r chi.Router) {
	//	r.Use(auth.BasicAuthMiddleware)
	//	r.Get("/history", handlers.HistoryHandler)
	//})

	//start server
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
