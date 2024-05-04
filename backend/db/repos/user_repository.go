package repos

import (
	"example.com/m/v2/db/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
    ListUsers() ([]*models.User, error)
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