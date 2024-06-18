package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	database "mmr/backend/db"
	"mmr/backend/db/repos"
	view "mmr/backend/models"
	"net/http"
)

type ProfileController struct{}

// GetProfile godoc
//
//	@Summary		Get profile
//	@Description	Get profile details of the authenticated user
//	@Tags 			Profile
//	@Produce		json
//	@Success		200	{object}	view.ProfileDetails
//	@Router			/v1/profile [get]
func (uc ProfileController) GetProfile(c *gin.Context) {
	// Initialize user repository
	userRepo := repos.NewUserRepository(database.DB)

	// Fetch user by identity user ID
	identityUserID := c.GetString("identityUserId")
	user, err := userRepo.GetByIdentityUserId(identityUserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	// Return user as JSON response
	c.JSON(http.StatusOK, view.ProfileDetailsViewFromModel(user))
}

// ClaimUser godoc
//
//	@Summary		Claim user
//	@Description	Claims a user by ID
//	@Tags 			Profile
//	@Param			user	body		view.ClaimUser	true	"User to be claimed"
//	@Produce		json
//	@Success		200	{object}	view.ProfileDetails
//	@Router			/v1/profile/claim [post]
func (uc ProfileController) ClaimUser(c *gin.Context) {
	var json view.ClaimUser
	err := c.ShouldBind(&json)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initialize user repository
	userRepo := repos.NewUserRepository(database.DB)

	// Check if identity has already claimed a user
	identityUserID := c.GetString("identityUserId")
	if identityUserID == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claimedUser, err := userRepo.GetByIdentityUserId(identityUserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	if claimedUser != nil && claimedUser.ID != 0 {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "User already claimed"})
		return
	}

	// Fetch user by ID
	user, err := userRepo.ClaimUserByID(json.UserID, identityUserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Return user as JSON response
	c.JSON(http.StatusOK, view.ProfileDetailsViewFromModel(user))
}
