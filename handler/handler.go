package handler

import (
	"golang-latihan-users-api/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := h.userService.GetUsers()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "400",
			"message": err,
		})
		return
	}

	var usersRsps []user.UserResponse

	for _, user := range users {
		userRsps := userResponseHandler(user)

		usersRsps = append(usersRsps, userRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Semua data pengguna telah berhasil ditampilkan",
		"data":    usersRsps,
	})
}

func (h *userHandler) GetOneUserHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user, err := h.userService.GetUser(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": err,
		})
		return
	}

	oneUserResponse := userResponseHandler(user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Data id tersebut telah berhasil ditampilkan!",
		"data":    oneUserResponse,
	})
}

func (h *userHandler) CreateNewUserHandler(c *gin.Context) {
	var userInput user.UserInput

	err := c.ShouldBindJSON(&userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": err,
		})
		return
	}

	createUser, err := h.userService.CreateUser(userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": err,
		})
		return
	}

	createUserResponse := userResponseHandler(createUser)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Data user baru telah berhasil ditampilkan",
		"data":    createUserResponse,
	})
}

func (h *userHandler) UpdateExistingUserHandler(c *gin.Context) {
	var userInput user.UserInput
	err := c.ShouldBindJSON(&userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user, err := h.userService.UpdateUser(id, userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": err.Error(),
		})
		return
	}

	updateUserResponse := userResponseHandler(user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Data pada id tersebut sudah berhasil diupdate!",
		"data":    updateUserResponse,
	})
}

func (h *userHandler) DeleteExistingUserHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user, err := h.userService.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": err.Error(),
		})
	}

	deleteUserResponse := userResponseHandler(user)

	updateMessageJSON := "Data id " + idString + " ini sudah berhasil dihapus!"

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": updateMessageJSON,
		"data":    deleteUserResponse,
	})
}

func userResponseHandler(userRsps user.User) user.UserResponse {
	return user.UserResponse{
		UserID:          userRsps.UserId,
		UserName:        userRsps.UserName,
		UserEmail:       userRsps.UserEmail,
		UserAddress:     userRsps.UserAddress,
		UserPhoneNumber: userRsps.UserPhoneNumber,
	}
}
