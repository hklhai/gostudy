package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"xorm.io/xorm"
)

var engine *xorm.Engine

type TbUser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func main() {

	engine, err := xorm.NewEngine("mysql",
		"root:mysql@tcp(127.0.0.1:3307)/didi?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("db connect error", err)
	}

	var u TbUser
	has, err := engine.Where("name=?", "xiaozhao").Get(&u)
	if err != nil {
		fmt.Println(err)
	}

	if has {
		s := strconv.Itoa(u.Id) + ":" + u.Name + ":" + u.Desc
		fmt.Println(s)
	}

}
