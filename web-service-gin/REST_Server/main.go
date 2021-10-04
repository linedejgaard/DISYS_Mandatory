package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", postCourse)
	router.DELETE("/courses/:id",deleteCourseByID)

	router.Run("localhost:8080")
}

