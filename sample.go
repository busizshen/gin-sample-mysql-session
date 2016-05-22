package smain

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	mysqlsession "github.com/kimiazhu/ginweb-contrib/sessions"
)

func main() {
	store, err := mysqlsession.NewMySQLStore(
		"root:@tcp(localhost:3306)/gin_sample?parseTime=true&loc=Local",
		"sessionstore",
		[]byte("secret-key"),
	)

	if err != nil {
		panic("failed to create session store")
	}
	store.Options(sessions.Options{
		Path: "/",
		//MaxAge: 300,
	})

	r := gin.Default()
	r.Use(sessions.Sessions("session", store))
	r.GET("/", func(c *gin.Context) {
		s = sessions.Default(c)
	})
	r.Run() // listen and server on 0.0.0.0:8080
}
