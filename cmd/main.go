package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zoro/pkg/config"
	"zoro/pkg/db"
	"zoro/pkg/redis"
	"zoro/pkg/routes"
	services "zoro/pkg/service"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.InitConnectDB(c)
	redisClient := redis.GetInstance(c)

	sv := services.MyServer{
		H:   h,
		Cof: c,
		R:   redisClient,
	}

	r := gin.Default()
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	routes.AdSunRoutes(r, &sv)
	r.Run(c.Port)

}
