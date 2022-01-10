package main

import (
	"api/api/rest"
	"api/pkg/config"
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

// @host 127.0.0.1:3200
// @BasePath /api/v1
func main() {
	// e-ticaret-api.herokuapp.com
	rest.RestRun(config.GetEnvironment("PORT", ".env"))
}
