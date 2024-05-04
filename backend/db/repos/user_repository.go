package repos

import (
	"mmr/backend/db/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
    ListUsers() ([]*models.User, error)
    FindUserByName(name string) (*models.User, error)
    SaveUser(user *models.User) error
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

func (ur *UserRepository) FindUserByName(name string) (*models.User, error) {
    user := &models.User{}
    if err := ur.db.Where("name = ?", name).First(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (ur *UserRepository) SaveUser(user *models.User) error {
    if err := ur.db.Save(user).Error; err != nil {
        return err
    }
    return nil
}