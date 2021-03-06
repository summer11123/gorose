package main

import (
	"github.com/gohouse/gorose/examples/config"
	"fmt"
	"github.com/gohouse/gorose"
	_ "github.com/lib/pq"
)

func main() {
	connection, err := gorose.Open(config.DbConfig, "postgres_dev")
	if err != nil {
		fmt.Println(err)
		return
	}
	// close DB
	defer connection.Close()

	db := connection.NewDB()

	//res := db.Table("users").First()
	//fmt.Println(res)

	// return json
	res2, err := db.Table("users").Limit(2).Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(db.LastSql)
	fmt.Println(res2)

}
