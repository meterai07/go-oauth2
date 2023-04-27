package handlers

import (
	"belajar-oauth/config"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GoogleLogin(c *gin.Context) {
	googleConfig := config.Config()
	url := googleConfig.AuthCodeURL(os.Getenv("AUTHCODEURL"))

	c.Redirect(http.StatusSeeOther, url)
}

func GoogleCallback(c *gin.Context) {
	state := c.Request.URL.Query().Get("state") // state is used to prevent CSRF, read more about it here: https://auth0.com/docs/protocols/state-parameters
	if state != os.Getenv("AUTHCODEURL") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	code := c.Request.URL.Query().Get("code") // the actual code to get the token

	googleConfig := config.Config() // get the google config

	token, err := googleConfig.Exchange(c, code) // get the token from the code
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken) // get the user data
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userData, err := io.ReadAll(resp.Body) // read the response
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.String(http.StatusOK, string(userData))
}
