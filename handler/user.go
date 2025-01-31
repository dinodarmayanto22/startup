package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/startup/helper"
	"github.com/startup/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	// Bind JSON input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Invalid request payload", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return // Pastikan untuk mengembalikan setelah mengirim respons
	}

	// Register user
	user, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Failed to register user", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return // Pastikan untuk mengembalikan setelah mengirim respons
	}

	// Jika berhasil, kirim respons sukses
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}
