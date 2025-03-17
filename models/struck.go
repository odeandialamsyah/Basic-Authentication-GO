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

type Role struct {
    gorm.Model
    Name  string `json:"name" gorm:"unique;not null"`
    Users []User `gorm:"foreignKey:RoleID" json:"-"`
}

func MigrateUsers(db *gorm.DB) {
    db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
}