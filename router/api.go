package router

import (
	"github.com/gin-gonic/gin"
	"gin-redis-manager/app/Http/controller"
)

func Api (r *gin.Engine) {
	r.POST("/testConnect", controller.Home.TestConnect)
}