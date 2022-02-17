package ent

import (
	config2 "api/pkg/config"
	_ "github.com/lib/pq"
)

type Database struct {
	DriverName     string
	DataSourceName string
}

func EntConnection() *Client {
	cfg := config2.GetConf()
	pg := Database{
		DriverName:     cfg.Database.DriverName,
		DataSourceName: cfg.Database.DataSourceName,
	}
	client, _ := Open(pg.DriverName, pg.DataSourceName)

	return client
}
