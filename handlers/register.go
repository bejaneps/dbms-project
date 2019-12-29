package handlers

import (
	"fmt"
	"net/http"

	"github.com/bejaneps/dbms-project/crud"

	"github.com/gin-gonic/gin"
)

var registerSuccess = "user registered successfully, id: "

// RegisterHandler is a handler for /register path, which serves register.html page
func RegisterHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

// RegisterSubmitHandler is a handler for processing request from /register/submit page
func RegisterSubmitHandler(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/register")
	}

	password := c.Query("password")
	if password == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/register")
	}

	uType := c.Query("type")
	if uType == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/register")
	}

	id := crud.AddUser(name, password, uType)

	//check if registered before for cookies
	if val := c.Query("type"); val != "" {
		c.SetCookie("type", "", 86400, "", "127.0.0.1", false, false)
	}
	if val := c.Query("id"); val != "" {
		c.SetCookie("id", "", 86400, "", "127.0.0.1", false, false)
	}

	if uType == "std" {
		c.SetCookie("type", "std", 86400, "", "127.0.0.1", false, false)
		c.HTML(http.StatusOK, "register.html", fmt.Sprintf(registerSuccess+"%d", id))
	} else if uType == "tcr" {
		c.SetCookie("type", "tcr", 86400, "", "127.0.0.1", false, false)
		c.HTML(http.StatusOK, "register.html", fmt.Sprintf(registerSuccess+"%d", id))
	} else if uType == "cdr" {
		c.SetCookie("type", "cdr", 86400, "", "127.0.0.1", false, false)
		c.HTML(http.StatusOK, "register.html", fmt.Sprintf(registerSuccess+"%d", id))
	}
}
