package repos

import (
	"errors"
	"mmr/backend/db/models"

	"github.com/mafredri/go-trueskill"
	"gorm.io/gorm"
)

type IUserRepository interface {
	ListUsers() ([]*models.User, error)
	GetOrCreateByName(name string) (*models.User, error)
	SaveUser(user *models.User) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) ListUsers() ([]*models.User, error) {
	var users []*models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetOrCreateByName(name string) (*models.User, error) {
	// Attempt to find an existing user by name
	user := &models.User{}
	if err := ur.db.Where("name = ?", name).First(user).Error; err != nil {
		// User not found in the database, let's create a new user
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newUser := &models.User{Name: name, MMR: 0, Mu: trueskill.DefaultMu, Sigma: trueskill.DefaultSigma}
			if err := ur.db.Create(newUser).Error; err != nil {
				return nil, err // Return error if unable to create user
			}
			return newUser, nil // Return the newly created user
		}
		return nil, err // Return other errors encountered during database operation
	}

	return user, nil // Return the found user
}

func (ur *UserRepository) SaveUser(user *models.User) (*models.User, error) {
	if err := ur.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
