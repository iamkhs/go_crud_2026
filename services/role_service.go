package services

import (
	"go_crud_2026/dto/response"
	"go_crud_2026/models"
	"go_crud_2026/repositories"
)

type RoleService struct {
	roleRepo *repositories.RoleRepository
}

func NewRoleService(roleRepo *repositories.RoleRepository) *RoleService {
	return &RoleService{roleRepo: roleRepo}
}

func (s *RoleService) GetAllRoles() []response.RoleResponse {
	roles := s.roleRepo.GetAllRoles()
	var roleResponseList []response.RoleResponse

	for _, role := range roles {
		roleResponse := response.RoleResponse{
			Id:          role.ID,
			Role:        role.Name,
			Description: role.Description,
		}
		roleResponseList = append(roleResponseList, roleResponse)
	}
	
	return roleResponseList
}

func (s *RoleService) CreateRole(role models.Role) (models.Role, bool) {
	return s.roleRepo.Create(role)
}
