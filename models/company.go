package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model

	CompanyName string `json:"company_name" binding:"required"`
	Country     string `json:"country" binding:"required"`
	City        string `json:"city" binding:"nullable"`
	Address     string `json:"address" binding:"required"`
	State       string `json:"state" binding:"nullable"`
	Zip         string `json:"zip" binding:"nullable"`
	Phone       string `json:"phone" binding:"required"`
	CompanyLogo string `json:"company_logo" binding:"nullable"`
	UserId      uint   `json:"user_id" binding:"required"`
}
