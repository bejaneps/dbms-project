package handlers

import (
	"net/http"
	"strconv"

	"github.com/bejaneps/dbms-project/crud"
	"github.com/bejaneps/dbms-project/models"
	"github.com/gin-gonic/gin"
)

// CdrCourseHandler is a handler for 'add course' button for coordinator
func CdrCourseHandler(c *gin.Context) {
	var err error

	course := models.Course{
		CourseID: c.Query("course_id"),
		Title:    c.Query("title"),
		DeptName: c.Query("department"),
	}
	course.Credits, err = strconv.Atoi(c.Query("credits"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	err = crud.AddCourse(course)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
}

// CdrTeacherHandler is a handler for 'add teacher' button for coordinator
func CdrTeacherHandler(c *gin.Context) {
	var err error

	teacher := models.Instructor{
		Name:     c.Query("name"),
		DeptName: c.Query("department"),
		Salary:   c.Query("salary"),
	}

	err = crud.AddTeacher(teacher)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
}

// CdrStudentHandler is a handler for 'add student' button for coordinator
func CdrStudentHandler(c *gin.Context) {
	var err error

	student := models.Student{
		Name:     c.Query("name"),
		DeptName: c.Query("department"),
	}
	student.TotCred, err = strconv.Atoi(c.Query("credits"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	err = crud.AddStudent(student)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
}

// CdrAssignCourseHandler is a handler for 'assign course to instructor' button for coordinator
func CdrAssignCourseHandler(c *gin.Context) {
	var err error

	teaches := models.Teaches{
		CourseID: c.Query("course_id"),
	}
	teaches.IID, err = strconv.Atoi(c.Query("i_id"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	err = crud.AssignCourse(teaches)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
}

// CdrAssignInstructorHandler is a handler for 'assign instructor to student' button for coordinator
func CdrAssignInstructorHandler(c *gin.Context) {
	var err error

	advisor := models.Advisor{}
	advisor.SID, err = strconv.Atoi(c.Query("s_id"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	advisor.IID, err = strconv.Atoi(c.Query("i_id"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	err = crud.AssignInstructor(advisor)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
}
