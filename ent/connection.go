package ent

import (
	config2 "api/pkg/config"
	"context"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	DriverName     string
	DataSourceName string
}

func ConnectionEnt() *Client {
	cfg := config2.GetConf()
	pg := Database{
		DriverName:     cfg.Database.DriverName,
		DataSourceName: cfg.Database.DataSourceName,
	}
	client, _ := Open(pg.DriverName, pg.DataSourceName)
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Println(err.Error())
	}
	return client
}
