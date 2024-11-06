package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64          `gorm:"<-:create;primaryKey"`
	UserName  string         `gorm:"<-:create;unique"`
	Password  string         `gorm:"not null"`
	NickName  sql.NullString `gorm:"<-:create"`
	Email     sql.NullString `gorm:"<-:create;unique"`
	Avatar    sql.NullString
	Status    bool `gorm:"default:true"`
	IsAdmin   bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
