package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/startup/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 	dsn := "host=127.0.0.1 user=dino password=dino1234 dbname=startup_crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	fmt.Println("Connection to database is good")

	// 	var users []user.User
	// 	length := len(users)

	// 	fmt.Println(length)

	// 	db.Find(&users)

	// 	length = len(users)
	// 	fmt.Println(length)

	// 	for _, user := range users {
	// 		fmt.Println(user.Name)
	// 		fmt.Println(user.Email)
	// 		fmt.Println("=================")

	// 	}
	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
}

func handler(c *gin.Context) {
	dsn := "host=127.0.0.1 user=dino password=dino1234 dbname=startup_crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
