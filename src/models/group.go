package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name      string
	GroupCode string
	Users     []UserGroup
}
