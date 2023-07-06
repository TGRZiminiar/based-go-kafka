package repositories

import "time"

type User struct {
	ID        string `gorm:"type:uuid;primary_key;not null"`
	UserName  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Role      string `gorm:"default:user;not null"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
