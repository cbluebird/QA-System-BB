package router

import "github.com/gin-gonic/gin"

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("login")
	}
}
