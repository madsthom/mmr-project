package models

import _ "ariga.io/atlas-provider-gorm/gormschema"
import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"unique"`
	MMR   int
	Mu    float64
	Sigma float64
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
	TeamOneID uint
	TeamOne   Team `gorm:"foreignKey:TeamOneID"`
	TeamTwoID uint
	TeamTwo   Team `gorm:"foreignKey:TeamTwoID"`
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
