package db

import (
	"time"
)

type User struct {
	Id        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Todos     []Todo `gorm:"many2many:todos_users"`
}
type Todo struct {
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
