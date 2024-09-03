package models

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
)
import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `gorm:"unique"`
	DisplayName    *string
	MMR            int     // TODO: Migrate usage of this field to player history
	Mu             float64 // TODO: Migrate usage of this field to player history
	Sigma          float64 // TODO: Migrate usage of this field to player history
	IdentityUserId *string `gorm:"unique"`
}

type PlayerHistory struct {
	gorm.Model
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
	MMR     int
	Mu      float64
	Sigma   float64
	MatchID uint
	Match   Match `gorm:"foreignKey:MatchID"`
}

type Team struct {
	gorm.Model
	UserOneID uint
	UserOne   User `gorm:"foreignKey:UserOneID"`
	UserTwoID uint
	UserTwo   User `gorm:"foreignKey:UserTwoID"`
	Score     uint
	Winner    bool
}

type Match struct {
	gorm.Model
	TeamOneID       uint
	TeamOne         Team `gorm:"foreignKey:TeamOneID"`
	TeamTwoID       uint
	TeamTwo         Team            `gorm:"foreignKey:TeamTwoID"`
	MMRCalculations *MMRCalculation `gorm:"foreignKey:MatchID"`
	SeasonID        uint
	Season          Season `gorm:"foreignKey:SeasonID"`
}

type MMRCalculation struct {
	gorm.Model
	MatchID                  uint
	Match                    Match `gorm:"foreignKey:MatchID"`
	TeamOnePlayerOneMMRDelta int
	TeamOnePlayerTwoMMRDelta int
	TeamTwoPlayerOneMMRDelta int
	TeamTwoPlayerTwoMMRDelta int
}

type Season struct {
	gorm.Model
}
