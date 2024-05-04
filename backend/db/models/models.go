package models

import "gorm.io/gorm"

type User struct {
  gorm.Model
  Name string
  MMR  int
}

type Team struct {
  gorm.Model
  UserOneID uint
  UserOne   User `gorm:"foreignKey:UserOneID"`
  UserTwoID uint
  UserTwo   User `gorm:"foreignKey:UserTwoID"`
}

type Match struct {
  gorm.Model
  TeamOneID uint
  TeamOne   Team `gorm:"foreignKey:TeamOneID"`
  TeamTwoID uint
  TeamTwo   Team `gorm:"foreignKey:TeamTwoID"`
}