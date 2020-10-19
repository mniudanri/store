package main

import (
	"github.com/mniudanri/store/config"
	"github.com/mniudanri/store/router"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	setter := config.GetConfig()
	defer setter.DB.Close()

	router.InitRoute(":5100")
}
