package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name      string
	GroupCode int
	Users     []UserGroup
}
