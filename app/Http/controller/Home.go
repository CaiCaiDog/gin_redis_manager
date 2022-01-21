package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
	"github.com/go-redis/redis"
	//"gin-redis-manager/app/Http/model"
)

// index
type HomeController struct {}

var (
	Home = &HomeController {}
)

// 首页
func (i *HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

// 测试连接
func (I *HomeController) TestConnect(c *gin.Context) {
	// host
	host := c.DefaultPostForm("host", "")
	if host == "" {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "host不能为空"})
	}

	// host
	port := c.DefaultPostForm("port", "")
	if port == "" {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "host不能为空"})
	}

	// auth
	pass := c.DefaultPostForm("auth", "")

	// 建立redis连接
	con := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
		Password: pass,
	})
	
	// 连接测试
	_, err := con.Ping().Result()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err})
	}
	c.JSON(http.StatusOK, gin.H{"code": 200})
}