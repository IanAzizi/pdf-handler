//handler/userhandler
package handler

import (
	"net/http"
	constants "pdf-APP/const"
	"pdf-APP/model"
	"pdf-APP/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

var Users []model.UserModel

func (u *UserHandler) GetAllUser(c *gin.Context) {
	v, ok := c.Get("user")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var current *model.UserModel
	switch t := v.(type) {
	case *model.UserModel:
		current = t
	case model.UserModel:
		current = &t
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid user in context"})
		return
	}

	if current.Role != constants.Manager {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	all := services.GetAllUsers()
	c.JSON(http.StatusOK, all)
}
func UserRegister(c *gin.Context) {

}
func (u *UserHandler) UserRegister(c *gin.Context) {
	var NewUser model.UserModel
	if err := c.ShouldBindJSON(&NewUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	create, err := services.RegisterService(NewUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, create)
}

func(u *UserHandler) LoginHandler(c *gin.Context) {
	 var creds struct {
        Id       string `json:"id"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, ok := services.Authenticate(creds.Id, creds.Password)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
        return
    }

    c.JSON(http.StatusOK, user)
}
