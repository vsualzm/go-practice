package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/cheking-api", CheckingAPI)

	// port
	router.Run(":8080")
}

func CheckingAPI(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API RUNNING",
	})
}
