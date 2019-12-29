package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogoutHandler is a handler for /logout path, which cleans all user cookies.
func LogoutHandler(c *gin.Context) {
	c.SetCookie("type", "", 86400, "", "127.0.0.1", false, false)
	c.SetCookie("id", "", 86400, "", "127.0.0.1", false, false)

	c.Redirect(http.StatusTemporaryRedirect, "/login")
}
