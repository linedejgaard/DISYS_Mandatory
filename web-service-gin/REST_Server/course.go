package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// course represents data about a record course.
type course struct {
	ID            string  `json:"id"`
	StudentRating float64 `json:"studentrating"`
	Name          string  `json:"name"`
	Teacher       string  `json:"teacher"`
	ECTS          float64 `json:"ects"`
}

// courses slice to seed record course data.
var courses = []course{
	{ID: "1", StudentRating: 9.1, Name: "GRPRO", Teacher: "Dan", ECTS: 15},
	{ID: "2", StudentRating: 8.6, Name: "BPAK", Teacher: "Henriette", ECTS: 7.5},
	{ID: "3", StudentRating: 8.1, Name: "BFOC", Teacher: "Allesandro", ECTS: 7.5},
	{ID: "4", StudentRating: 7.8, Name: "BSAD", Teacher: "Thore", ECTS: 7.5},
	{ID: "5", StudentRating: 8.3, Name: "BFYP", Teacher: "Troels", ECTS: 15},
	{ID: "6", StudentRating: 6.2, Name: "SYSDAB", Teacher: "Louise", ECTS: 7.5},
	{ID: "7", StudentRating: 7.8, Name: "BDSA", Teacher: "Paolo", ECTS: 15},
	{ID: "8", StudentRating: 8.3, Name: "DISYS", Teacher: "Marco", ECTS: 7.5},
	{ID: "9", StudentRating: 6.1, Name: "IDBS", Teacher: "Bjórn", ECTS: 7.5},
}

// postCourse adds a course from JSON received in the request body.
func postCourse(c *gin.Context) {
	var newCourse course

	// Call BindJSON to bind the received JSON to
	// newCourse.
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	// Add the new course to the slice.
	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

// getCourseByID locates the course whose ID value matches the id
// parameter sent by the client, then returns that course as a response.
func getCourseByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of course, looking for
	// a course whose ID value matches the parameter.
	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}

// getCourses responds with the list of all courses as JSON.
func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func deleteCourseByID(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	courses = RemoveIndex(courses, id - 1)
	c.IndentedJSON(http.StatusCreated, courses)
}

func RemoveIndex(s []course, index int) []course {
    return append(s[:index], s[index+1:]...)
}
