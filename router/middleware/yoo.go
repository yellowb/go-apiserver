package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Yoo(c *gin.Context) {
	log.Println("Yooooooooo......")
	c.Next()
}
