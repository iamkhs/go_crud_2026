package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password"`
	Enable   bool   `json:"enable"`
	Otp      string `json:"otp"`

	CompanyID *uint
	Company   Company `gorm:"foreignKey:CompanyID"`
}
