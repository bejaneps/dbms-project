package handlers

import (
	"net/http"
	"strconv"

	"github.com/bejaneps/dbms-project/crud"
	"github.com/gin-gonic/gin"
)

// StdListCoursesHandler is a handler for 'see available courses' button for student
func StdListCoursesHandler(c *gin.Context) {
	sID, _ := c.Cookie("id")
	studentID, err := strconv.Atoi(sID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	courses, err := crud.ListStdCourses(studentID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	if len(courses) == 0 {
		c.HTML(http.StatusOK, "dashboard_std.html", nil)
		return
	}

	c.HTML(http.StatusOK, "dashboard_std.html", courses)
	return
}
