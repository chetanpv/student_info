package main

import (
	"studentDetails/gomodule/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// My router configuration with rest calls
	router := gin.Default()

	router.GET("/student", routes.GetStudentDetails)
	router.POST("/student", routes.PostStudentDetails)
	router.GET("/student/:id", routes.GetStudentDetails)

	// Connect to Database
	// database.ConnectToDB()

	// Starting the GIN server
	router.Run()
}
