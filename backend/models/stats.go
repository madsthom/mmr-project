package view

import (
	"mmr/backend/db/models"
	"time"
)

type PlayerHistoryDetails struct {
	UserID uint      `json:"userId" binding:"required"`
	Name   string    `json:"name" binding:"required"`
	Date   time.Time `json:"date" binding:"required"`
	MMR    int       `json:"mmr" binding:"required"`
}

func PlayerHistoryDetailsViewFromModel(playerHistory models.PlayerHistory) PlayerHistoryDetails {
	return PlayerHistoryDetails{
		UserID: playerHistory.UserID,
		Name:   playerHistory.User.Name,
		Date:   playerHistory.Match.CreatedAt,
		MMR:    playerHistory.MMR,
	}
}

type TimeStatisticsEntry struct {
	// 0-6, 0 is Sunday
	DayOfWeek int `json:"dayOfWeek" binding:"required"`
	// 0-23
	HourOfDay int `json:"hourOfDay" binding:"required"`
	Count     int `json:"count" binding:"required"`
}
