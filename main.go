package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "hello golang!",
	})
}

func main(){
	r := gin.Default()//返回默认的路由引擎
	//指定用户使用GET请求访问/hello时， 执行sayHello这个函数
	r.GET("/hello", sayHello)

	//r.GET("/book", ...)
	//r.POST("/create_book", ...)
	//r.POST("/update_book", ...)
	//r.POST("/delete_book", ...)

	//RESTful 架构API
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	//启动服务
	r.Run(":9090")//默认8080端口
}