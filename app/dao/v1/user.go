package dao_v1

import (
	"fuge/app/core"
	models "fuge/app/models/v1"
)

type userDAO struct {
}

var UserDAO *userDAO = &userDAO{}

func (s *userDAO) DoGetUserByOpenID(openID string) *models.User {
	user := &models.User{
		OpenID: openID,
	}
	if err := core.DB.Where("open_id = ?", user.OpenID).First(user).Error; err != nil {
		core.DB.Create(user)
	}

	return user
}
