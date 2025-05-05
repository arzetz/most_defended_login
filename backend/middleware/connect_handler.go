package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func ConnectHandler(c *gin.Context, db *sql.DB) {
	c.JSON(200, gin.H{"message": "База данных подключена!"})
}
