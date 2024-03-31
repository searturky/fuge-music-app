package service_v1

import (
	daos "fuge/app/dao/v1"
	models "fuge/app/models/v1"
)

type wechatService struct {
}

var WechatService *wechatService = &wechatService{}

func (s *wechatService) LoginWechat(lwi *models.LoginWeChatIn) {
	data := daos.WechatDAO.DoLoginWechat(lwi)
	ensureUser(data)
	// return data
}

func ensureUser(res *models.LoginWeChatRes) {
	OpenID := res.OpenID
	user := daos.UserDAO.DoGetUserByOpenID(OpenID)
	print(&user)
	return
}
