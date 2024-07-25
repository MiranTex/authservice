package src

import "github.com/gin-gonic/gin"

func StartRoute() {

	r := gin.Default()

	userpassroutes := r.Group("userpassmodule")
	{
		userpassroutes.GET("/register")
	}
	// startRouting(r);

}
