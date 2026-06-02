package main

import (
	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/v1/healthcheck", app.HealthcheckHandler)
	router.POST("/blog", app.CreateBlog)
	router.GET("/blog/:Id", app.GetBlogById)
	return router
}
