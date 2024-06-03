package view

import "mmr/backend/db/models"

type UserDetails struct {
	UserID      uint    `json:"userId" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	DisplayName *string `json:"displayName"`
}

func UserDetailsViewFromModel(user models.User) UserDetails {
	return UserDetails{
		UserID:      user.ID,
		Name:        user.Name,
		DisplayName: user.DisplayName,
	}
}
