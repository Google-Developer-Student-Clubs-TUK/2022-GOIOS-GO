package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Password string
	Groups   []UserGroup
}

type UserGroup struct {
	gorm.Model
	UserID        uint
	User          User `gorm:"foreignKey:UserID;references:ID"`
	GroupID       uint
	Group         Group         `gorm:"foreignKey:GroupID;references:ID"`
	AvailableTime pq.Int64Array `gorm:"type:integer[]"`
}

//func (user UserGroup) BeforeSave(db *gorm.DB) error {
//	user.CreateAt = time.Now()
//	return nil
//}