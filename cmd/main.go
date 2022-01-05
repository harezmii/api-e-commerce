package main

import (
	"e-commerce-api/api/rest"
	"e-commerce-api/pkg/config"
)

// @title E Commerce API
// @version 1.0
// @description This is e-commerce server.
// @termsOfService https://e-ticaret-api.herokuapp.com/api/v1/

// @contact.name Api Support
// @contact.url https://e-ticaret-api.herokuapp.com/support
// @contact.email suatcnby06@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host e-ticaret-api.herokuapp.com
// @BasePath /api/v1
func main() {
	rest.RestRun(config.GetEnvironment("PORT", ".env"))
}
