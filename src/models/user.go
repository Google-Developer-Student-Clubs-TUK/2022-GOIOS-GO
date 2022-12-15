package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"uniqueIndex"`
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
	Master        string
}

//func (user UserGroup) BeforeSave(db *gorm.DB) error {
//	user.CreateAt = time.Now()
//	return nil
//}
