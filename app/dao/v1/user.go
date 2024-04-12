package dao_v1

import (
	"fuge/app/core"
	models "fuge/app/models/v1"
)

type userDAO struct {
}

var UserDAO *userDAO = &userDAO{}

func (s *userDAO) DoGetUserByOpenID(openID string) (*models.User, error) {
	var user models.User
	err := core.DB.Where("open_id = ?", openID).First(&user).Error
	return &user, err
}

func (s *userDAO) DoCreateUser(openID string) *models.User {
	user := &models.User{
		OpenID: openID,
		Status: models.Uncomplete,
	}
	core.DB.Create(user)
	return user
}

func (s *userDAO) DoGetUserByUserID(userID int) (*models.User, error) {
	var user models.User
	err := core.DB.Where("id = ?", userID).First(&user).Error
	return &user, err
}
