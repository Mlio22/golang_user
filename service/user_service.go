package service

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	dtos "golang_user/DTOs"
	"golang_user/models"
)

type UserService interface {
	GetUsers(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserServiceStr struct {
}

func NewUserService() UserService {
	return &UserServiceStr{}
}

var users = []models.User{}

// nextUserID is a simple global counter user ids.
var nextUserID uint = 1

func (s *UserServiceStr) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (s *UserServiceStr) GetUser(c *gin.Context) {
	userID, ok := parseUserID(c)
	if !ok {
		return
	}

	index := findUserIndex(userID)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users[index]})
}

func (s *UserServiceStr) CreateUser(c *gin.Context) {
	input, ok := bindUserInput(c)
	if !ok {
		return
	}

	// check if email already exists in the user list,
	// if exists return 409
	if emailExists(input.Email, 0) {
		c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		return
	}

	user := models.User{
		ID:        nextUserID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}
	users = append(users, user)
	nextUserID++

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// update user by id
func (s *UserServiceStr) UpdateUser(c *gin.Context) {
	selectedUserId, ok := parseUserID(c)
	if !ok {
		return
	}

	input, ok := bindUserInput(c)
	if !ok {
		return
	}

	// if user not found, return 404
	index := findUserIndex(selectedUserId)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// if email already exists for another user, return 409
	if emailExists(input.Email, selectedUserId) {
		c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		return
	}

	user := &users[index]
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email

	c.JSON(http.StatusOK, gin.H{"data": users[index]})
}

func (s *UserServiceStr) DeleteUser(c *gin.Context) {
	userID, ok := parseUserID(c)
	if !ok {
		return
	}

	index := findUserIndex(userID)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	users = append(users[:index], users[index+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

// bind user input into DTO with validation.
func bindUserInput(c *gin.Context) (dtos.UserInput, bool) {
	var input dtos.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return dtos.UserInput{}, false
	}

	input.FirstName = strings.TrimSpace(input.FirstName)
	input.LastName = strings.TrimSpace(input.LastName)
	input.Email = strings.TrimSpace(strings.ToLower(input.Email))
	if input.FirstName == "" || input.LastName == "" || input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "first_name, last_name, and email are required"})
		return dtos.UserInput{}, false
	}

	return input, true
}

// get user id from request context.
// validate if it's positive
func parseUserID(c *gin.Context) (uint, bool) {
	parsedID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || parsedID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return 0, false
	}

	return uint(parsedID), true
}

// func for get user by index.
// it needed on separate func to avoid code duplication
func findUserIndex(userID uint) int {
	for index, user := range users {
		if user.ID == userID {
			return index
		}
	}

	return -1
}

// check if user list on memory list
func emailExists(email string, currentUserId uint) bool {
	for _, user := range users {
		if user.ID != currentUserId && user.Email == email {
			return true
		}
	}

	return false

}
