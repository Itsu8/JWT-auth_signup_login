package router

import (
	"github.com/Itsu8/Auth/controllers"
	"github.com/gin-gonic/gin"
	"github.com/Itsu8/Auth/middleware"
)

func setHttpQueries(r *gin.Engine) {
	r.POST("/signup", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
}

func RunServer() {
	router := gin.Default()
	setHttpQueries(router)
	defer router.Run()
}
