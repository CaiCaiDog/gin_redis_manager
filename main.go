package main

import (
	"gin-redis-manager/bootstrap"
)

func main() {
	// 加载服务
	bootstrap.InitServer()
}