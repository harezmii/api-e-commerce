package ent

import (
	"api/ent"
	_ "github.com/go-sql-driver/mysql"
)

var ClientEnt = entConnection()

func entConnection() *Client {
	client, _ := ent.Open("mysql", "root:@tcp(localhost:3306)/deneme?parseTime=True")
	return client
}
func Entdisconnection() {
	defer ClientEnt.Close()
}
