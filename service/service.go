package service

import (
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserService interface { 
	AddUser(*gin.Context)
	GetUser(*gin.Context)
}

type User struct { 
	ID   int64  `json:"id"`
	Nombre string `json:"n"`
}

type userService struct { 
	data map[int64]User
}

func (us *userService) AddUser(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = time.Now().Unix()
	us.data[user.ID] = user
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (us *userService) GetUser(c *gin.Context) { 
	c.JSON(http.StatusCreated, us.data)
}

func NewUserService() UserService { 
	m := make(map[int64]User)
	return &userService{data: m}
}
