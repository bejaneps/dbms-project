package handlers

import (
	"net/http"
	"strconv"

	"github.com/bejaneps/dbms-project/crud"

	"github.com/gin-gonic/gin"
)

// TcrListCoursesHandler is a handler for 'see assigned courses' button, for teacher
func TcrListCoursesHandler(c *gin.Context) {
	iID, _ := c.Cookie("id")
	instructorID, err := strconv.Atoi(iID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	courses, err := crud.ListTcrCourses(instructorID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	if len(courses) == 0 {
		c.HTML(http.StatusOK, "dashboard_tcr.html", nil)
		return
	}

	c.HTML(http.StatusOK, "dashboard_tcr.html", courses)
	return
}
