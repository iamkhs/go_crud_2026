package repositories

import (
	"go_crud_2026/database"
	"go_crud_2026/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository() *RoleRepository {
	err := database.DB.AutoMigrate(models.Role{})
	if err != nil {
		panic("Failed to migrate Role model!")
	}
	return &RoleRepository{db: database.DB}
}

func (r *RoleRepository) GetAllRoles() []models.Role {
	var roles []models.Role
	r.db.Find(&roles)
	return roles
}

func (r *RoleRepository) Create(role models.Role) (models.Role, bool) {
	result := r.db.Create(&role)
	if result.Error != nil {
		return models.Role{}, false
	}
	return role, true
}
