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
	StoreRanking(matchID uint, userID uint, mu float64, sigma float64, mmr int) (*models.PlayerHistory, error)
	StoreMatchMMRCalculation(matchID uint, player1Delta int, player2Delta int, player3Delta int, player4Delta int) (*models.MMRCalculation, error)
	GetLatestPlayerHistory(playerID uint) (*models.PlayerHistory, error)
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

func (ur *UserRepository) StoreRanking(matchID uint, userID uint, mu float64, sigma float64, mmr int) (*models.PlayerHistory, error) {
	playerHistory := &models.PlayerHistory{
		UserID:  userID,
		MMR:     mmr,
		Mu:      mu,
		Sigma:   sigma,
		MatchID: matchID,
	}
	if err := ur.db.Create(playerHistory).Error; err != nil {
		return nil, err
	}
	return playerHistory, nil
}

func (ur *UserRepository) StoreMatchMMRCalculation(matchID uint, player1Delta int, player2Delta int, player3Delta int, player4Delta int) (*models.MMRCalculation, error) {
	mmrCalculation := &models.MMRCalculation{
		MatchID:                  matchID,
		TeamOnePlayerOneMMRDelta: player1Delta,
		TeamOnePlayerTwoMMRDelta: player2Delta,
		TeamTwoPlayerOneMMRDelta: player3Delta,
		TeamTwoPlayerTwoMMRDelta: player4Delta,
	}
	if err := ur.db.Create(mmrCalculation).Error; err != nil {
		return nil, err
	}
	return mmrCalculation, nil
}

func (ur *UserRepository) GetLatestPlayerHistory(playerID uint) (*models.PlayerHistory, error) {
	playerHistory := &models.PlayerHistory{}
	if err := ur.db.Where("user_id = ?", playerID).Order("created_at desc").First(playerHistory).Error; err != nil {
		return nil, err
	}
	return playerHistory, nil
}
