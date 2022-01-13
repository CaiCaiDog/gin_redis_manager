package router

import (
	"github.com/gin-gonic/gin"
	"gin-redis-manager/app/Http/controller"
)

// 
func Web (r *gin.Engine) {
	r.GET("/", controller.Index.Index)
	r.GET("/welcome", controller.Index.Welcome)
}