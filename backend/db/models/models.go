package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"unique"`
	MMR   int
	Mu    float64
	Sigma float64
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
