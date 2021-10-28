package controller

import (
	"github.com/chengziwj/mudan/domain/entity"
	"github.com/chengziwj/mudan/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	newUser, err := s.us.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (s *UserController) GetUser(c *gin.Context) {
	uid, _ := strconv.ParseUint(c.Param("uid"), 10, 64)
	user, err := s.us.GetUser(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
