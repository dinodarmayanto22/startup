package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/startup/handler"
	"github.com/startup/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*************  ✨ Codeium Command 🌟  *************/

func main() {
	dsn := "host=127.0.0.1 user=dino password=dino1234 dbname=startup_crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewSevice(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()

}
