package server

import (
	// "net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	// router.GET("/test", getHello)
	router.Run("localhost:8080")

}

// func getHello(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, "{hi: 'hello'}")
// }
