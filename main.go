package main

import (
	root "./api/v1/root"
	mysession "./api/v1/session"
	"./web"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	msession "github.com/kimiazhu/ginweb-contrib/sessions"
)

func main() {
	store, err := msession.NewMySQLStore(
		"root:@tcp(localhost:3306)/gin_sample?parseTime=true&loc=Local",
		"sessionstore",
		[]byte("secret-key"),
	)

	if err != nil {
		panic("failed to create session store")
	}
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 300,
	})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Use(sessions.Sessions("session", store))
	r.GET("/", web.GET)
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/", root.GET)
		v1.POST("/", root.POST)
		v1.PUT("/", root.PUT)
		v1.DELETE("/", root.DELETE)
	}
	session := v1.Group("/session")
	{
		session.GET("/", mysession.GET)
		session.DELETE("/", mysession.DELETE)

	}
	r.Run() // listen and server on 0.0.0.0:8080
}
