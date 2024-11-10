package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64          `gorm:"<-:create;primaryKey" json:"ID"`
	UserName  string         `gorm:"<-:create;unique" json:"UserName"`
	Password  string         `gorm:"not null" json:"Password"`
	NickName  sql.NullString `gorm:"<-:create" json:"NickName"`
	Email     sql.NullString `gorm:"<-:create;unique" json:"Email"`
	Avatar    sql.NullString `json:"Avatar"`
	Status    bool           `gorm:"default:true" json:"Status"`
	IsAdmin   bool           `gorm:"default:false" json:"IsAdmin"`
	CreatedAt time.Time      `gorm:"<-:create" json:"CreatedAt"`
	UpdatedAt time.Time      `gorm:"<-:update" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"DeletedAt"`
}
