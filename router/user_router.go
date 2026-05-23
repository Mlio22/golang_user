package router

import (
	"github.com/gin-gonic/gin"

	dtos "golang_user/DTOs"
	"golang_user/service"
)

var userService = service.NewUserService()

// Router-level aliases keep Swagger schemas near route docs while reusing DTO definitions.
type UserInputDTO = dtos.UserInput
type UserResponseDTO = dtos.UserResponse
type UsersResponseDTO = dtos.UsersResponse
type ErrorResponseDTO = dtos.ErrorResponse
type MessageResponseDTO = dtos.MessageResponse

// UserRouter registers the user CRUD endpoints under /api/users.
func UserRouter(r *gin.Engine) {
	users := r.Group("/api/users")

	users.GET("", getUsers)
	users.GET(":id", getUser)
	users.POST("", createUser)
	users.PUT(":id", updateUser)
	users.DELETE(":id", deleteUser)
}

// getUsers godoc
// @Summary List users
// @Description Returns all users
// @Tags users
// @Produce json
// @Success 200 {object} UsersResponseDTO
// @Router /api/users [get]
func getUsers(c *gin.Context) {
	userService.GetUsers(c)
}

// getUser godoc
// @Summary Get user
// @Description Returns one user by ID.
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} UserResponseDTO
// @Failure 400 {object} ErrorResponseDTO
// @Failure 404 {object} ErrorResponseDTO
// @Router /api/users/{id} [get]
func getUser(c *gin.Context) {
	userService.GetUser(c)
}

// createUser godoc
// @Summary Create user
// @Description Creates a user. and increase id counter
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserInputDTO true "User payload"
// @Success 201 {object} UserResponseDTO
// @Failure 400 {object} ErrorResponseDTO
// @Failure 409 {object} ErrorResponseDTO
// @Router /api/users [post]
func createUser(c *gin.Context) {
	userService.CreateUser(c)
}

// updateUser godoc
// @Summary Update user
// @Description Updates user by id.
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body UserInputDTO true "User payload"
// @Success 200 {object} UserResponseDTO
// @Failure 400 {object} ErrorResponseDTO
// @Failure 404 {object} ErrorResponseDTO
// @Failure 409 {object} ErrorResponseDTO
// @Router /api/users/{id} [put]
func updateUser(c *gin.Context) {
	userService.UpdateUser(c)
}

// deleteUser godoc
// @Summary Delete user
// @Description Deletes one user from the in-memory list.
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} MessageResponseDTO
// @Failure 400 {object} ErrorResponseDTO
// @Failure 404 {object} ErrorResponseDTO
// @Router /api/users/{id} [delete]
func deleteUser(c *gin.Context) {
	userService.DeleteUser(c)
}
