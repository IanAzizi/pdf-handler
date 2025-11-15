// router/userrouter
package router

import (
	"pdf-APP/api/handler"
	"pdf-APP/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	handler := handler.NewUserHandler()

	r.POST("/login", handler.LoginHandler)
	r.POST("/register", handler.UserRegister)

	protected := r.Group("/")
	protected.Use(middleware.UserMiddleware)
	{
		protected.GET("/getUsers", handler.GetAllUser)
	}
}
