package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	logger "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

type IdTable struct {
	Id int64
}

func main() {
	db_name := os.Getenv("db_name")
	var err error
	db, err := xorm.NewEngine("mysql", db_name)
	if err != nil {
		fmt.Printf("db1 error is %v", err)
		return
	}
	fmt.Printf("db_name: %v\n", db_name)
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		id := Get_id(db, db_name)
		c.JSON(200, gin.H{
			"id": id,
		})
	})
	go func() {
		logger.Println(http.ListenAndServe(":6060", nil))
	}()

	go func() {
		if err := http.ListenAndServe(":8888", nil); err != nil {
			logger.Error("initPprofHttpAsync error", zap.Error(err))
		}
	}()
	r.Run()
}

func Get_id(db *xorm.Engine, db_name string) (ID int64) {
	id := new(IdTable)
	cnt, err := db.Insert(id)
	if err != nil {
		fmt.Printf("The db1 error is %v", err)
		return
	}
	fmt.Printf("cnt: %d, id: %d\n", cnt, id.Id)
	return id.Id
}
