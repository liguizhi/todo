package pg

import (
	"github.com/gin-gonic/gin"
	"guizhi.me/my-todo/model"
)

//Todolist new
func Todolist(c *gin.Context) {
	msg := model.Getone()
	c.JSON(200, gin.H{"code": 200, "msg": msg})
}

//GetDb new
func GetDb(c *gin.Context) {
	msg := model.GetFromDb(c)
	c.JSON(200, gin.H{"code": 200, "msg": msg})
}

//SaveItem savedate
func SaveItem(c *gin.Context) {
	bindId := c.Request.FormValue("base_id")
	date := c.Request.FormValue("date")
	count := c.Request.FormValue("count")
	msg := map[string]interface{}{"bindId": bindId, "date": date, "count": count}
	c.JSON(200, gin.H{"code": 200, "input": msg})
}
