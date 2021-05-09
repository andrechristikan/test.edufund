package main

import (
	"github.com/gin-gonic/gin"

	"log"
	"os"

	"test.edufund/controllers"
	"test.edufund/db"

	"github.com/joho/godotenv"
)

//CORSMiddleware ...
//CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}


func main() {
	// Start
	r := gin.Default()

	//Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}

	// Middleware
	r.Use(CORSMiddleware())

	//Start PostgreSQL database
	//Example: db.GetDB() - More info in the models folder
	db.Init()

	// Versioning
	v1 := r.Group("/v1")
	{
		/*** START API ***/
		loanController := new(controllers.LoanController)
		v1.GET("/loans", loanController.All)
		v1.POST("/loans", loanController.Create)
		v1.PATCH("/loans/update/:id", loanController.Update)
		v1.PATCH("/loans/approve/:id", loanController.Approve)
	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test endpoints",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "not found",
		})
	})


	//Set Environment
	port := os.Getenv("APP_PORT")

	if os.Getenv("APP_ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Run(":"+port)
}