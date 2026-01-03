package services

import (
	"fmt"
	"go_crud_2026/auth"
	"go_crud_2026/models"
	"go_crud_2026/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() []models.User {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetById(id int) (*models.User, bool) {
	return s.repo.GetUserById(id)
}

func (s *UserService) CreateUser(user models.User) models.User {
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(id int, user models.User) (*models.User, bool) {
	return s.repo.Update(id, user)
}

func (s *UserService) DeleteUser(id int) bool {
	return s.repo.Delete(id)
}

func (s *UserService) Login(email, password string) (string, bool) {
	user, found := s.repo.GetByEmail(email)
	fmt.Println(user)
	if !found {
		return "", false
	}

	// For demo: no password hashing yet
	// In real app, compare hashed password
	if password != "password" { // Demo password
		return "", false
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", false
	}

	return token, true
}
