package database

import (
	"fmt"
	"golang-latihan-users-api/handler"
	"golang-latihan-users-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var NewDatabase *gorm.DB

func DBConnect() {
	dbc := "root:root@tcp(127.0.0.1:3306)/db_golang_latihan_users_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbc), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database is succesfully connected!")

	// db.Migrator().DropTable(&user.User{})

	db.AutoMigrate(&user.User{})

	usersRepository := user.NewRepository(db)
	usersService := user.NewService(usersRepository)
	usersHandler := handler.NewUserHandler(usersService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/users", usersHandler.GetAllUsersHandler)
	v1.GET("/users/:id", usersHandler.GetOneUserHandler)
	v1.POST("/users", usersHandler.CreateNewUserHandler)
	v1.PATCH("/users/:id", usersHandler.UpdateExistingUserHandler)
	v1.DELETE("/users/:id", usersHandler.DeleteExistingUserHandler)

	// v1.GET("/", user.HelloHandler)

	router.Run(":3600")
}
