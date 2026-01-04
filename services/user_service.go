package services

import (
	"go_crud_2026/auth"
	"go_crud_2026/dto/response"
	"go_crud_2026/models"
	"go_crud_2026/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() []response.UserResponse {
	users := s.repo.GetAllUsers()
	var userResponseList []response.UserResponse
	for _, user := range users {
		userResponse := response.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		userResponseList = append(userResponseList, userResponse)
	}

	return userResponseList
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
	if !found {
		return "", false
	}

	// Compare provided password with stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", false
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", false
	}

	return token, true
}
