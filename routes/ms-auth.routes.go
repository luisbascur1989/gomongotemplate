package routes

import (
	"comunty/ms-auth/api"

	"github.com/gin-gonic/gin"
)

func Auth(r *gin.RouterGroup) {
	r.POST("/api/register", api.RegisterHandler)
	r.POST("/api/login", api.RegisterHandler)
	r.GET("/api/login", api.RegisterHandler)
	r.GET("/api/users", api.GetUserHandler)
}
