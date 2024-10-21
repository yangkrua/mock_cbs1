package routers

import (
	apiController "mock_cbs1/controllers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter function will perform all route operations
func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//status info
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET") //, PUT, DELETE, UPDATE
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Content-Security-Policy, X-Content-Type-Options, X-Frame-Options")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("X-Frame-Options", "SameOrigin")
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=3600; includeSubDomains")
		c.Writer.Header().Set("Cache-Control", "max-age=3600")
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self';")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}

	})

	r.POST("/api/rlcs/account_info", apiController.InquiryCBSData)
	r.POST("/api/rlcs/account_payment_record", apiController.Account_payment_record)
	r.POST("/api/rlcs/account_statement", apiController.Account_Statement)

	return r

}
