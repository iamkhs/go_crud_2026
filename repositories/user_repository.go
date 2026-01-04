package repositories

import (
	"errors"
	"fmt"
	"go_crud_2026/database"
	"go_crud_2026/models"
	"strings"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Error migrating User model:", err)
		panic("Failed to migrate User model")
	}
	return &UserRepository{db: database.DB}
}

func (r *UserRepository) GetAllUsers() []models.User {
	var users []models.User
	r.db.Find(&users)
	return users
}

func (r *UserRepository) GetUserById(id int) (*models.User, bool) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, false
	}
	return &user, true
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value") {
			return models.User{}, errors.New("email already exists")
		}
		return models.User{}, result.Error
	}
	return user, nil
}

// Update user
func (r *UserRepository) Update(id int, updated models.User) (*models.User, error) {
	updated.ID = uint(id)
	if err := r.db.Save(&updated).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// Delete user
func (r *UserRepository) Delete(id int) bool {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return false
	}
	return true
}

func (r *UserRepository) GetByEmail(email string) (*models.User, bool) {
	var user models.User
	fmt.Println("email", email)
	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, false
	}
	return &user, true
}
