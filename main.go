package main

import (
	"golangsinau/handler"
	"golangsinau/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// This is a function that connects to the database.
	dsn := "root:@tcp(127.0.0.1:3306)/golang_sinau?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// This is a conditional statement that checks if there is an error in the connection. If there is an
	// error, the program will stop and print the error message.
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	
	router.Run()
}
