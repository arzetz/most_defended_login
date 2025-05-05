package main

import (
	"fmt"
	"log"
	"main/db"
	mw "main/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/api/login", func(c *gin.Context) {
		fmt.Println("Ооо бля")
		c.Status(200)
	})

	r.PATCH("/api/connectdb", mw.ConnectHandler(c*gin.Context, database))
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	r.Run(":8080")
} // Дописать подруб дб по кнопке на фронте
