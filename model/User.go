//model/user
package model


type UserModel struct{
	Id string 	`json:"id"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Role string `json:"role"`
}