package controllers

import (
	"github.com/gin-gonic/gin"
	database "mmr/backend/db"
	"mmr/backend/db/repos"
	view "mmr/backend/models"
	"net/http"
	"strconv"
)

type UsersController struct{}

// ListUsers godoc
//
//	@Summary		List users
//	@Description	Lists all users
//	@Tags 			Users
//	@Produce		json
//	@Success		200	{object}	[]view.UserDetails
//	@Router			/v1/users [get]
func (uc UsersController) ListUsers(c *gin.Context) {
	// Initialize user repository
	userRepo := repos.NewUserRepository(database.DB)

	// Fetch all users
	users, err := userRepo.ListUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, []view.UserDetails{})
		return
	}

	var userDetails []view.UserDetails

	for _, user := range users {
		userDetails = append(userDetails, view.UserDetailsViewFromModel(*user))
	}

	c.JSON(http.StatusOK, userDetails)
}

// CreateUser godoc
//
//	@Summary		Create user
//	@Description	Creates a new user
//	@Tags 			Users
//	@Param			user	body	view.CreateUser	true	"User data"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	view.UserDetails
//	@Router			/v1/users [post]
func (uc UsersController) CreateUser(c *gin.Context) {
	var json view.CreateUser
	err := c.ShouldBind(&json)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initialize user repository
	userRepo := repos.NewUserRepository(database.DB)

	// Check if user already exists
	_, err = userRepo.GetByName(json.Name)
	if err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// Create user
	user, err := userRepo.CreateByName(json.Name, json.DisplayName)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, view.UserDetailsViewFromModel(*user))
}

// SearchUsers godoc
//
//	@Summary		Search users
//	@Description	Searches users by name
//	@Tags 			Users
//	@Param			query	query	string	true	"Name to search for"
//	@Produce		json
//	@Success		200	{object}	[]view.UserDetails
//	@Router			/v1/users/search [get]
func (uc UsersController) SearchUsers(c *gin.Context) {
	// Initialize user repository
	userRepo := repos.NewUserRepository(database.DB)

	// Parse query from request
	query := c.Query("query")

	// Fetch users by name
	users, err := userRepo.SearchUsers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, []view.UserDetails{})
		return
	}

	var userDetails []view.UserDetails
	for _, user := range users {
		userDetails = append(userDetails, view.UserDetailsViewFromModel(*user))
	}

	// Return users as JSON response
	c.JSON(http.StatusOK, userDetails)
}

// GetUser godoc
//
//	@Summary		Get user
//	@Description	Get user by ID
//	@Tags 			Users
//	@Param			id	path	int	true	"User ID"
//	@Produce		json
//	@Success		200	{object}	view.UserDetails
//	@Router			/v1/users/{id} [get]
func (uc UsersController) GetUser(c *gin.Context) {
	// Initialize user repository
	userRepo := repos.NewUserRepository(database.DB)

	// Parse user ID from request
	userIdString := c.Param("id")
	userId, err := strconv.ParseUint(userIdString, 10, 64)

	// Fetch user by ID
	user, err := userRepo.GetByID(uint(userId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	// Return user as JSON response
	c.JSON(http.StatusOK, view.UserDetailsViewFromModel(*user))
}
