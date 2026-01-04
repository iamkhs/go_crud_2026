package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model

	CompanyName  string `json:"company_name" binding:"required"`
	EmployeeSize int    `json:"employee_size" binding:"nullable"`
	Country      string `json:"country"`
	City         string `json:"city" binding:"nullable"`
	Address      string `json:"address"`
	State        string `json:"state" binding:"nullable"`
	Zip          string `json:"zip" binding:"nullable"`
	Phone        string `json:"phone"`
	CompanyLogo  string `json:"company_logo" binding:"nullable"`
	UserId       uint   `json:"user_id" binding:"required"`
}
