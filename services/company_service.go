package services

import (
	"go_crud_2026/models"
	"go_crud_2026/repositories"
)

type CompanyService struct {
	repo *repositories.CompanyRepository
}

func NewCompanyService(repo *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{
		repo: repo,
	}
}

func (s *CompanyService) CreateCompany(company models.Company) models.Company {
	return s.repo.Create(company)
}

func (s *CompanyService) UpdateCompany(id int, company models.Company) (*models.Company, bool) {
	return s.repo.Update(id, company)
}
