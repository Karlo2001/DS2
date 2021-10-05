package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type course struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	Satisfaction_Rating float32 `json:"satisfaction_rating"`
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.GET("/courses/:id", GetCourseById)
	router.Run("localhost:8080")
}

var courses = []course{
	{ID: "1", Name: "Distributed Systems", Satisfaction_Rating: 60.3},
	{ID: "2", Name: "First Year Project", Satisfaction_Rating: 80.9},
}

func GetCourseById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}
