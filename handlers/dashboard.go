package handlers

import (
	"net/http"
	"strconv"

	"github.com/bejaneps/dbms-project/crud"

	"github.com/gin-gonic/gin"
)

// DashboardHandler is a handler for /dashboard path, which redirects user to student, teacher or coordinator html pages.
func DashboardHandler(c *gin.Context) {
	uType, err := c.Cookie("type")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	id, err := c.Cookie("id")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	realID, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err.Error())
	}

	_, err = crud.CheckUser(realID, "")
	if err != nil {
		c.Redirect(http.StatusUnauthorized, "/login")
	}

	if uType == "std" {
		c.HTML(http.StatusOK, "dashboard_std.html", nil)
	} else if uType == "tcr" {
		c.HTML(http.StatusOK, "dashboard_tcr.html", nil)
	} else if uType == "cdr" {
		c.HTML(http.StatusOK, "dashboard_cdr.html", nil)
	} else {
		c.HTML(http.StatusOK, "error.html", "unauthorized")
	}
}
