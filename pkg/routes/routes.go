package routes

import (
	"github.com/gin-gonic/gin"
	services "zoro/pkg/service"
)

func AdSunRoutes(r *gin.Engine, sv *services.MyServer) {
	routes := r.Group("/api/test")
	routes.GET("/get_info", sv.GetDb)
	routes.GET("/get_all", sv.GetAllDb)
}
