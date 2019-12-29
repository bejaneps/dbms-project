package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bejaneps/dbms-project/crud"

	"github.com/gin-gonic/gin"
)

var errIDNotNumber = errors.New("id should be a number")

// IndexHandler is a handler for / path, which redirects to /login page
func IndexHandler(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/login")
}

// LoginHandler is a handler for /login path, which server login.html page
func LoginHandler(c *gin.Context) {
	_, err := c.Cookie("type")
	if err == nil {
		if val, _ := c.Cookie("id"); val != "" {
			c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
		}
	}

	c.HTML(http.StatusOK, "login.html", nil)
}

// LoginSubmitHandler is a handler for processing request /login/submit path
func LoginSubmitHandler(c *gin.Context) {
	var err error

	_, err = c.Cookie("type")
	if err == nil {
		if val, _ := c.Cookie("id"); val != "" {
			c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
		}
	}

	id := c.Query("id")
	if id == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	realID, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", errIDNotNumber)
	}

	password := c.Query("password")
	if password == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	uType := c.Query("type")
	if uType == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	_, err = crud.CheckUser(realID, password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", err.Error())
	}

	c.SetCookie("type", uType, 86400, "", "127.0.0.1", false, false)
	c.SetCookie("id", id, 86400, "", "127.0.0.1", false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
}
