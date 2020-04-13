package td

import (
	"github.com/gin-gonic/gin"
	"guizhi.me/my-todo/model"
)

//Todolist new
func Todolist(c *gin.Context) {
	msg := model.Getone()
	c.JSON(200, gin.H{"code": 200, "msg": msg})
}

//Todolist new
func GetDb(c *gin.Context) {
	msg := model.GetFromDb(c)
	c.JSON(200, gin.H{"code": 200, "msg": msg})
}
