package Router

import (
	"authservice/src/userpassmodule"

	"github.com/gin-gonic/gin"
)

func StartRoute() *gin.Engine {

	r := gin.Default()

	userpassroutes := r.Group("userpassmodule")
	{
		userpassroutes.GET("/register", userpassmodule.Register)
	}

	return r
}
