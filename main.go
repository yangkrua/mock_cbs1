package main

import (
	_ "mock_cbs1/docs"
	Routers "mock_cbs1/routers"
)

// @title		MOCK RLCS Service
// @version		1.0
// @description	Testing Swagger APIs.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			127.0.0.1:9343
// @BasePath
// @schemes		http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	router := Routers.SetupRouter()
	router.Run("localhost:9343")
}
