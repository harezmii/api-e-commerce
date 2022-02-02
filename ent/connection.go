package ent

import (
	_ "github.com/go-sql-driver/mysql"
)

func EntConnection() *Client {
	client, _ := Open("mysql", "root:@tcp(localhost:3306)/deneme?parseTime=True")
	return client
}
