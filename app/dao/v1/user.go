package dao_v1

import (
	"fuge/app/core"
	models "fuge/app/models/v1"
)

type userDAO struct {
}

var UserDAO *userDAO = &userDAO{}

func (s *userDAO) DoGetUserByOpenID(openID string) (*models.User, error) {
	user := &models.User{
		OpenID: openID,
	}
	err := core.DB.Where("open_id = ?", user.OpenID).First(user).Error

	return user, err
}

func (s *userDAO) DoCreateUser(user *models.User) *models.User {
	core.DB.Create(user)
	return user
}
