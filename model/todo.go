package model

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

//MyItem local type
type MyItem struct {
	Name string `json:"jname" form:"go"`
	Age  int    `json:"jage" form:"to"`
}

//PullLog type
type PullLog struct {
	ID             int    `json:"id" form:"id"`
	BaseID         string `json:"baseId" form:"base_id"`
	Date           string `json:"date" form:"date"`
	Count          int    `json:"count" form:"count"`
	WriteCount     int    `json:"writeCount" form:"write_count"`
	Status         int    `json:"status" form:"status"`
	FetchCount     int    `json:"fetchCount" form:"fetch_count"`
	LastFetchTime  string `json:"lastFetchTime" form:"last_fetch_time"`
	CreateTime     string `json:"createTime" form:"create_time"`
	LastModifyTime string `json:"lastModifyTime" form:"last_modify_time"`
}

//user:password@type(ip:host)/dbname
var dbdsn = ""

//Getone []MyItem
func Getone() MyItem {
	persons := make([]MyItem, 0)
	res := MyItem{Name: "", Age: 0}
	res1 := MyItem{Name: "", Age: 0}
	res.Name = "agui"
	res.Age = 30
	persons = append(persons, res)
	res1.Name = "new"
	res1.Age = 20
	persons = append(persons, res1)
	return res
}

//GetFromDb return db item
func GetFromDb(c *gin.Context) []PullLog {
	db, err := sql.Open("mysql", dbdsn)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	rows, err := db.Query("SELECT id,base_id,date,count,last_fetch_time,create_time,last_modify_time,write_count,fetch_count,status FROM dx_data_pull_log limit ?", 10)
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	pulls := make([]PullLog, 0)

	for rows.Next() {
		var pulllog PullLog
		// rows.Scan(&pulllog)
		rows.Scan(&pulllog.ID, &pulllog.BaseID, &pulllog.Date, &pulllog.Count, &pulllog.LastFetchTime, &pulllog.CreateTime, &pulllog.LastModifyTime, &pulllog.WriteCount, &pulllog.FetchCount, &pulllog.Status)
		pulls = append(pulls, pulllog)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
	return pulls
}

func SaveItem(item PullLog) bool {
	db, err := sql.Open("mysql", dbdsn)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	db.ExecContext()
}
