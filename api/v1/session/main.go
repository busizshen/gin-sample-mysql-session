package session

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count += 1
	}
	session.Set("count", count)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func DELETE(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("count")
	session.Save()
	v := session.Get("count")
	if v != nil {
		c.JSON(http.StatusOK, gin.H{"message": "failed to delete session"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success to clear session"})
	}

}
