package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BohdanNikolaychuk/backend-learn/internal/core/config"
	"github.com/BohdanNikolaychuk/backend-learn/internal/core/db"
	authRep "github.com/BohdanNikolaychuk/backend-learn/internal/features/auth/repository"
	authService "github.com/BohdanNikolaychuk/backend-learn/internal/features/auth/service"
	authTransport "github.com/BohdanNikolaychuk/backend-learn/internal/features/auth/transport"
	userRep "github.com/BohdanNikolaychuk/backend-learn/internal/features/user/repository"
	userService "github.com/BohdanNikolaychuk/backend-learn/internal/features/user/service"
	userTransport "github.com/BohdanNikolaychuk/backend-learn/internal/features/user/transport"
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

	userRep := userRep.NewUserRepository(pool)
	userService := userService.NewUserService(userRep)
	userTransport := userTransport.NewUserTransport(userService)
	authRep := authRep.NewAuthRepository(pool)
	authService := authService.NewAuthService(userRep, authRep, conf.JWTSecret)
	authTransport := authTransport.NewAuthTransport(authService)

	fmt.Printf("server started on port: %s\n", conf.ServerPort)

	http.HandleFunc("POST /user", userTransport.Create)

	http.HandleFunc("GET /user/{id}", userTransport.GetByID)

	http.HandleFunc("DELETE /user/{id}", userTransport.DeleteByID)

	http.HandleFunc("PATCH /user/{id}", userTransport.Update)

	http.HandleFunc("PATCH /user/{id}/password", userTransport.ChangePassword)

	// auth
	http.HandleFunc("POST /auth/login", authTransport.Login)
	http.HandleFunc("POST /auth/refresh", authTransport.Refresh)
	http.HandleFunc("POST /auth/logout", authTransport.Logout)

	http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), nil)

}
