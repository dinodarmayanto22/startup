package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/startup/handler"
	"github.com/startup/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=127.0.0.1 user=dino password=dino1234 dbname=startup_crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewSevice(userRepository)
	userByEmail, err := userRepository.FindByEmail("dinodarmayanto22@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(userByEmail)

	if userByEmail.ID == 0 {
		fmt.Println("user tidak di temukan")
	} else {
		fmt.Println(userByEmail.Name)
	}
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()

	// user := user.User{
	// 	Name: "Test simpan",
	// }

	// userRepository.Save(user)
}
