package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username	string `json:"username" gorm:"unique;not null"`
	Email		string `json:"email" gorm:"unique;not null"`
	Password 	string `json:"password" gorm:"not null"`
	RoleID   uint   `json:"role_id" gorm:"not null"`
    Role     Role   `json:"role" gorm:"foreignKey:RoleID"`
}

func MigrateUsers(db *gorm.DB) {
    db.AutoMigrate(&User{})
}