package main

import (
	"github.com/Itsu8/Auth/initializers"
	"github.com/Itsu8/Auth/modules"
	"github.com/Itsu8/Auth/router"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&modules.User{})
}

func main(){
	router.RunServer()
}