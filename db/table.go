package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Todos     []Todo `gorm:"many2many:users_todos"`
}
type Todo struct {
	gorm.Model
	Id        uint   `gorm:"prmaryKey"`
	Title     string `gorm:"not null"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Deadline    string
	Tag         string
	TagColor    string
	CreaterId   uint `gorm:"not null"`
	RepeatFlag  bool `gorm:"default:false"`
	RepeatSpan  uint
	RepeatCount uint
}
