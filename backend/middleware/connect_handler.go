package middleware

import (
	"database/sql"
	"log"
	"main/db"

	"github.com/gin-gonic/gin"
)

var GlobalDB *sql.DB

func ConnectHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		if GlobalDB != nil {
			c.JSON(400, gin.H{"message": "БД уже подключена"})
			return
		}
		database, err := db.Connect()
		if err != nil {
			log.Fatal(err)
			return
		} else {
			GlobalDB = database
			c.JSON(200, gin.H{"message": "База данных подключена!"})
		}
	}
}

func DisconnectHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		if GlobalDB == nil {
			c.JSON(400, gin.H{"message": "База не подключена"})
			return
		}
		if err := GlobalDB.Close(); err != nil {
			log.Println("ошибка при закрытии базы: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		GlobalDB = nil
		c.JSON(200, gin.H{"message": "БД успешно закрыта"})
	}
}
