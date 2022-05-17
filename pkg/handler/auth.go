package handler

import (
	"net/http"

	"github.com/AsaHero/todolist/db/models"
	"github.com/gin-gonic/gin"
)

type signinUser struct {
	Username string `json:"username"`
	Password_hash string `json:"password_hash"`
}

func (h Handler) SignIn(c *gin.Context) {
	var input signinUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token ,err := h.service.Authorization.GenerateToken(input.Username, input.Password_hash)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":token,
	})
} 

func (h Handler) SignUp(c *gin.Context) {
	var input models.Users

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id ,err := h.service.Authorization.CreateAccount(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})
}
