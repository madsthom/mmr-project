package repos

import (
	"database/sql"
	"errors"
	"mmr/backend/db/models"

	"github.com/mafredri/go-trueskill"
	"gorm.io/gorm"
)

type IUserRepository interface {
	ListUsers() ([]*models.User, error)
	GetOrCreateByName(name string) (*models.User, error)
	GetByName(name string) (*models.User, error)
	CreateByName(name string, displayName *string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByIdentityUserId(identityUserId string) (*models.User, error)
	SearchUsers(query string) ([]*models.User, error)
	SaveUser(user *models.User) (*models.User, error)
	StoreRanking(matchID uint, userID uint, mu float64, sigma float64, mmr int) (*models.PlayerHistory, error)
	StoreMatchMMRCalculation(matchID uint, player1Delta int, player2Delta int, player3Delta int, player4Delta int) (*models.MMRCalculation, error)
	GetLatestPlayerHistory(playerID uint) (*models.PlayerHistory, error)
	ListPlayerHistory(playerID *uint) ([]*models.PlayerHistory, error)
	ClearPlayerHistories()
	ClaimUserByID(userID uint, identityUserID string) (*models.User, error)
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
	user, err := ur.GetByName(name)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ur.CreateByName(name, nil)
		}
		return nil, err // Return other errors encountered during database operation
	}

	return user, nil // Return the found user
}

func (ur *UserRepository) GetByName(name string) (*models.User, error) {
	// Attempt to find an existing user by name
	user := &models.User{}
	err := ur.db.Where("name = ?", name).First(user).Error
	return user, err
}

func (ur *UserRepository) CreateByName(name string, displayName *string) (*models.User, error) {
	newUser := &models.User{Name: name, DisplayName: displayName, MMR: 0, Mu: trueskill.DefaultMu, Sigma: 2}
	if err := ur.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (ur *UserRepository) GetByID(id uint) (*models.User, error) {
	user := &models.User{}
	if err := ur.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetByIdentityUserId(identityUserId string) (*models.User, error) {
	user := &models.User{}
	if err := ur.db.Where("identity_user_id = ?", identityUserId).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) SearchUsers(query string) ([]*models.User, error) {
	var users []*models.User
	err := ur.db.
		Where("name ilike @query OR display_name ilike @query", sql.Named("query", "%"+query+"%")).
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
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

func (ur *UserRepository) ListPlayerHistory(playerID *uint) ([]*models.PlayerHistory, error) {
	var playerHistories []*models.PlayerHistory
	query := ur.db.Model(&models.PlayerHistory{}).
		Preload("User").
		Joins("Match").
		Order("\"Match\".created_at asc")

	if playerID != nil {
		query = query.Where("user_id = ?", *playerID)
	}

	err := query.Find(&playerHistories).Error
	if err != nil {
		return nil, err
	}
	return playerHistories, nil
}

func (ur *UserRepository) ClearPlayerHistories() {
	ur.db.Exec("TRUNCATE TABLE player_histories RESTART IDENTITY;")
}

func (ur *UserRepository) ClaimUserByID(userID uint, identityUserID string) (*models.User, error) {
	user := &models.User{}
	if err := ur.db.Where("id = ?", userID).First(user).Error; err != nil {
		return nil, err
	}

	user.IdentityUserId = &identityUserID
	if err := ur.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
