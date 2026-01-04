package repositories

import (
	"go_crud_2026/database"
	"go_crud_2026/models"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository() *CompanyRepository {
	err := database.DB.AutoMigrate(models.Company{})
	if err != nil {
		panic("Failed to migrate Company model!")
	}
	return &CompanyRepository{db: database.DB}
}

func (r *CompanyRepository) Create(company models.Company) models.Company {
	r.db.Create(&company)
	return company
}

func (r *CompanyRepository) Update(id int, company models.Company) (*models.Company, bool) {
	var existingCompany models.Company
	result := r.db.First(id)
	if result.Error != nil {
		return nil, false
	}
	existingCompany.CompanyName = company.CompanyName
	existingCompany.Country = company.Country
	existingCompany.City = company.City
	existingCompany.Address = company.Address
	existingCompany.State = company.State
	existingCompany.Zip = company.Zip
	existingCompany.Phone = company.Phone
	existingCompany.CompanyLogo = company.CompanyLogo

	r.db.Save(&existingCompany)
	return &existingCompany, true
}
