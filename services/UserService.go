// services/userservice
package services

import (
	"fmt"
	"pdf-APP/model"
	"sync"

	"github.com/gin-gonic/gin"
)

var users []model.UserModel
var UserMu sync.Mutex

func RegisterService(newUser model.UserModel) (model.UserModel, error) {
	UserMu.Lock()
	defer UserMu.Unlock()

	for _, u := range users {
		if u.Id == newUser.Id {
			return model.UserModel{}, fmt.Errorf("user with id %s already exists", newUser.Id)
		}
	}

	users = append(users, newUser)
	return newUser, nil
}

func LoginService(c *gin.Context) {

	id := c.Query("Id")
	password := c.Query("Password")
	UserMu.Lock()
	defer UserMu.Unlock()
	for _, user := range users {
		if user.Id == id && user.Password == password {
			c.JSON(200, user)
		} else {
			c.JSON(401, gin.H{"error": "Unauthorized"})
		}
	}
}
func GetAllUsers() []model.UserModel {
	UserMu.Lock()
	defer UserMu.Unlock()
	out := make([]model.UserModel, len(users))
	copy(out, users)

	return out
}
func Authenticate(id, password string) (*model.UserModel, bool) {
	UserMu.Lock()
	defer UserMu.Unlock()
	for _, u := range users {
		if u.Id == id && u.Password == password {
			UserCopy := u
			return &UserCopy, true
		}
	}
	return nil, false
}
