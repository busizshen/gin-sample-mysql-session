package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func POST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "POST"})
}

func PUT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "PUT"})
}

func DELETE(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "DELET"})
}
