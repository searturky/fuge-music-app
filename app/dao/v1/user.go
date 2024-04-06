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
	}
	core.DB.Create(user)
	return user
}
