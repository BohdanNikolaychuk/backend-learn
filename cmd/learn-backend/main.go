package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BohdanNikolaychuk/backend-learn/internal/core/config"
	"github.com/BohdanNikolaychuk/backend-learn/internal/core/db"
	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/repository"
	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/service"
	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/transport"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("conf connecting ")

	fmt.Println("db connecting ")
	pool, err := db.Connect(conf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db connected")

	userRep := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRep)
	userTransport := transport.NewUserTransport(userService)

	fmt.Printf("server started on port: %s\n", conf.ServerPort)

	http.HandleFunc("POST /user", userTransport.Create)

	http.HandleFunc("GET /user/{id}", userTransport.GetByID)

	http.HandleFunc("DELETE /user/{id}", userTransport.DeleteByID)

	http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), nil)

}
