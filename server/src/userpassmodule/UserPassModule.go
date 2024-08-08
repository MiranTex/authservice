package userpassmodule

import "github.com/gin-gonic/gin"

type requestParams struct {
	username string
	password string
}

func Register(c *gin.Context) {
	var params requestParams
	c.BindJSON(&params)
	c.JSON(200, gin.H{
		"message":  "Register",
		"username": params.username,
		"password": params.password,
	})
}
