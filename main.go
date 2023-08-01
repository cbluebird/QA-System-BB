package main

import (
	"QA-system/config/config"
	"QA-system/config/database"
	"QA-system/config/router"
	"QA-system/config/session"
	midware "QA-system/midware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.Init()
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midware.ErrHandler())
	r.NoMethod(midware.HandleNotFound)
	r.NoRoute(midware.HandleNotFound)
	session.Init(r)
	router.Init(r)

	err := r.Run(":" + config.Config.GetString("router.port"))
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
}
