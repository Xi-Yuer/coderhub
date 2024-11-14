package model

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"<-:create;primaryKey" json:"ID"`
	UserName  string         `gorm:"<-:create;unique" json:"UserName" validate:"required,min=3,max=32"`
	Password  string         `gorm:"not null" json:"Password" validate:"required,min=6,max=32"`
	NickName  sql.NullString `gorm:"<-:create" json:"NickName" validate:"required,min=3,max=32"`
	Email     sql.NullString `gorm:"<-:create;unique" json:"Email" validate:"required,email"`
	Avatar    sql.NullString `json:"Avatar"`
	Status    bool           `gorm:"default:true" json:"Status"`
	IsAdmin   bool           `gorm:"default:false" json:"IsAdmin"`
	CreatedAt time.Time      `gorm:"<-:create" json:"CreatedAt"`
	UpdatedAt time.Time      `gorm:"<-:update" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"DeletedAt"`
}

func (u *User) CacheKeyByID(id int64) string {
	return fmt.Sprintf("user:id:%d", id)
}

func (u *User) CacheKeyByName(name string) string {
	return fmt.Sprintf("user:name:%s", name)
}
