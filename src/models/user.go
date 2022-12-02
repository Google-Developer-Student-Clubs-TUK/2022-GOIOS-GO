package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName string
	Password string
	Groups   []*Group `gorm:"many2many:user_groups;"`
}

type UserGroup struct {
	UserID   int `gorm:"primaryKey"`
	GroupId  int `gorm:"primaryKey"`
	CreateAt time.Time
	DeleteAt gorm.DeletedAt
}

func (user UserGroup) BeforeSave(db *gorm.DB) error {
	user.CreateAt = time.Now()
	return nil
}
