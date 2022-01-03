package main

import (
	"e-commerce-api/api/rest"
	"e-commerce-api/pkg/config"
)

func main() {
	rest.RestRun(config.GetEnvironment("SERVER_PORT", ".env"))
}
