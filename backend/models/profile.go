package view

import "mmr/backend/db/models"

type ClaimUser struct {
	UserID uint `json:"userId" binding:"required"`
}

type ProfileDetails struct {
	UserID *uint `json:"userId"`
}

func ProfileDetailsViewFromModel(user *models.User) ProfileDetails {
	if user == nil {
		return ProfileDetails{}
	}

	return ProfileDetails{
		UserID: &user.ID,
	}
}
