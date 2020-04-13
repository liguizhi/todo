package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"guizhi.me/my-todo/controller/pg"
	"guizhi.me/my-todo/controller/td"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// 使用默认中间件创建一个gin路由器
	// logger and recovery (crash-free) 中间件
	router := gin.Default()

	router.GET("/firstget", getting)
	router.POST("/firstpost", posting)
	router.GET("/list", td.Todolist)
	router.GET("/getdb", td.GetDb)
	router.POST("/saveitem", pg.SaveItem)

	// router.PUT("/firstput", putting)
	// router.DELETE("/firstdelete", deleting)
	// router.PATCH("/firstpatch", patching)
	// router.HEAD("/firsthead", head)
	// router.OPTIONS("/firstoption", options)

	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run()
	// router.Run(":3000") for a hard coded port
}

func getting(c *gin.Context) {
	query := c.Query("id")
	result := gin.H{"code": 200, "msg": query}
	c.JSON(200, result)
}

func posting(c *gin.Context) {
	query := c.PostForm("name")
	result := gin.H{"code": 200, "msg": query}
	c.JSON(200, result)
}
