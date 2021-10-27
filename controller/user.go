package controller

import (
	"github.com/chengziwj/mudan/domain/entity"
	"github.com/chengziwj/mudan/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	us service.UserServiceInterface
}

func NewUserController(us service.UserServiceInterface) *UserController {
	return &UserController{us: us}
}

func (s *UserController) SaveUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	if err := user.BeforeSave(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	newUser, err := s.us.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newUser)
}
