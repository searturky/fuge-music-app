package service_v1

import (
	"fuge/app/core"
	daos "fuge/app/dao/v1"
	models "fuge/app/models/v1"
	"fuge/app/utils"
)

type wechatService struct {
}

var WechatService *wechatService = &wechatService{}

func (s *wechatService) LoginWechat(lwi *models.LoginWeChatIn) {
	data := daos.WechatDAO.DoLoginWechat(lwi)
	ensureUser(data)
	// return data
}

func ensureUser(res *models.LoginWeChatRes) *models.User {
	OpenID := res.OpenID
	user, err := daos.UserDAO.DoGetUserByOpenID(OpenID)
	if err != nil {
		user.Nickname = genRandomNickname()
		daos.UserDAO.DoCreateUser(user)
	}
	return user
}

func genRandomNickname() string {
	prefix := core.GetConf().DefaultUserPrefix
	randomName := utils.GetRandomName(prefix)
	return randomName
}
