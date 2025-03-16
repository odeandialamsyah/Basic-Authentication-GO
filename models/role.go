package models

import "gorm.io/gorm"

type Role struct {
    gorm.Model
    Name  string `json:"name" gorm:"unique;not null"`
    Users []User `gorm:"foreignKey:RoleID" json:"-"`
}

func MigrateRole(db *gorm.DB) {
    db.AutoMigrate(&Role{})
}