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
	pg := Database{
		DriverName:     config2.GetEnvironment("DRIVER_NAME", config2.STRING).(string),
		DataSourceName: config2.GetEnvironment("DATA_SOURCE_NAME", config2.STRING).(string),
	}
	client, _ := Open(pg.DriverName, pg.DataSourceName)
	return client
}
