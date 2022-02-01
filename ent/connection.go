package ent

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func EntConnection() *Client {
	client, _ := Open("mysql", "root:@tcp(localhost:3306)/deneme?parseTime=True")
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Println("Schema Error")
	}
	return client
}
