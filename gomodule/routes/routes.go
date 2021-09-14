package routes

import (
	"fmt"
	"net/http"
	"studentDetails/gomodule/database"
	"studentDetails/gomodule/models"

	"github.com/gin-gonic/gin"
)

var studentDetails = []models.StudentDetail{}

// PoststudentDetails adds studentDetails from request
func PostStudentDetails(c *gin.Context) {
	if err := c.BindJSON(&studentDetails); err != nil {
		return
	}
	fmt.Println("Student details: ", studentDetails)
	database.InsertIntoTable(studentDetails)
	c.IndentedJSON(http.StatusCreated, studentDetails)
}

// GetstudentDetails gets all student details
func GetStudentDetails(c *gin.Context) {
	result := database.GetAllFromTable()
	c.IndentedJSON(http.StatusOK, result)
}

// GetStudentDetailsID gets single student detail
func GetStudentDetailsID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of studentDetails, looking for
	// an studentDetail whose ID value matches the parameter.
	for _, a := range studentDetails {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "studentDetail not found"})
}
