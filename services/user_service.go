package services

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"go_crud_2026/auth"
	"go_crud_2026/dto/response"
	"go_crud_2026/models"
	"go_crud_2026/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo         *repositories.UserRepository
	emailService *EmailService
}

func NewUserService(repo *repositories.UserRepository, emailService *EmailService) *UserService {
	return &UserService{repo: repo, emailService: emailService}
}

func (s *UserService) GetAllUsers() []response.UserResponse {
	users := s.repo.GetAllUsers()
	var userResponseList []response.UserResponse
	for _, user := range users {
		userResponse := response.UserResponse{
			ID:    user.ID,
			Name:  user.FullName,
			Email: user.Email,
		}
		userResponseList = append(userResponseList, userResponse)
	}

	return userResponseList
}

func (s *UserService) GetById(id int) (*models.User, bool) {
	return s.repo.GetUserById(id)
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err == nil {
		user.Password = string(hashedPassword)
	}

	// Generate OTP
	otp := s.generateOTP()
	user.Otp = otp
	user.Enable = false

	// Save to DB
	createdUser, err := s.repo.Create(user)
	if err != nil {
		return models.User{}, err
	}
	
	// Send OTP Email
	go s.emailService.SendOtpEmail(createdUser.Email, otp)

	return createdUser, nil
}

func (s *UserService) generateOTP() string {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		fmt.Println("Error generating OTP:", err)
		return "123456" // Fallback
	}
	return fmt.Sprintf("%06d", n.Int64())
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
