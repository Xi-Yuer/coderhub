package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64   `gorm:"<-:create;primaryKey"`
	UserName  string  `gorm:"<-:create;unique"`
	Password  string  `gorm:"not null"`
	NickName  *string `gorm:"<-:create"`
	Email     *string `gorm:"<-:create;unique"`
	Avatar    *string
	Status    bool `gorm:"default:true"`
	IsAdmin   bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
