package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type IdTable struct {
	Id int64
}

func main() {

	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		id1, id2 := Get_ids("db1")
		c.JSON(200, gin.H{
			"id1": id1,
			"id2": id2,
		})
	})
	r.Run()
}

func Get_ids(db_name string) (ID1 int64, ID2 int64) {
	var err error
	db1, err := xorm.NewEngine("mysql", "root:test@tcp(127.0.0.1:3306)/id_service_db")
	if err != nil {
		fmt.Printf("db1 error is %v", err)
		return
	}
	db2, err := xorm.NewEngine("mysql", "root:test@tcp(127.0.0.1:3307)/id_service_db")
	if err != nil {
		fmt.Printf("db2 error is %v", err)
		return
	}
	id1 := new(IdTable)
	cnt1, err := db1.Insert(id1)
	if err != nil {
		fmt.Printf("The db1 error is %v", err)
		return
	}
	fmt.Printf("cnt1: %d, id1: %d\n", cnt1, id1.Id)
	id2 := new(IdTable)
	cnt2, err := db2.Insert(id2)
	if err != nil {
		fmt.Printf("The db2 error is %v", err)
		return
	}
	fmt.Printf("cnt2: %d, id2: %d\n", cnt2, id2.Id)
	return id1.Id, id2.Id
}
