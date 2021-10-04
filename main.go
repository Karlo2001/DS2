package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type course struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	Satisfaction_Rating float32 `json:"satisfaction_rating"`
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)

	router.Run("localhost:8080")
}

var courses = []course{
	{ID: 1, Name: "Distributed Systems", Satisfaction_Rating: 60.3},
	{ID: 2, Name: "First Year Project", Satisfaction_Rating: 80.9},
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}
