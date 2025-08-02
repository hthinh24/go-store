package main

import (
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	_ = logger.NewAppLogger("identity-service")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
