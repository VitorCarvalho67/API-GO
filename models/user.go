package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	u.Id = uuid.New().String()
// 	return
// }
