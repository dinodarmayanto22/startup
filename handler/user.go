/*************  ✨ Codeium Command 🌟  *************/
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
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Registerid account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return // Pastikan untuk mengembalikan setelah mengirim respons
	}
	// Register user
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Registerid account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return // Pastikan untuk mengembalikan setelah mengirim respons
	}

	formatter := user.FormatUser(newUser, "tokentokentokentoken")

	// Jika berhasil, kirim respons sukses
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {

}
