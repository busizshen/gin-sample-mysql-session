package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
