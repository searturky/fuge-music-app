package service_v1

import (
	// "context"
	"context"
	"encoding/json"
	"fuge/app/constant"
	"fuge/app/core"
	daos "fuge/app/dao/v1"
	models "fuge/app/models/v1"
	"fuge/app/utils"
	"time"
)

type wechatService struct {
}

var WechatService *wechatService = &wechatService{}

func (s *wechatService) GetPhoneNumber(code string, user *models.User) error {
	// 获取小程序Authtoken
	authToken, err := ensureWechatAuthToken()
	if err != nil {
		return err
	}
	// 获取手机号
	data, err := daos.WechatDAO.DoGetPhoneNumber(code, authToken)
	if err != nil {
		return err
	}
	err = daos.UserDAO.DoSavePhoneNumber(data, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *wechatService) LoginWechat(lwi *models.LoginWeChatIn) (map[string]interface{}, error) {
	data := daos.WechatDAO.DoLoginWechat(lwi)
	user := ensureUser(data)
	userSchema := &models.UserSchema{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Phone:     user.Phone,
		AvatarUrl: user.AvatarUrl,
		Status:    user.Status,
		Province:  user.Province,
		City:      user.City,
		Country:   user.Country,
		Gender:    user.Gender,
		Language:  user.Language,
	}
	v, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	userinfo := string(v)
	token := utils.GetRandomString(32)
	key := constant.UserToken + token
	setLoginTokenToRedis(key, userinfo)
	ret := make(map[string]interface{})
	ret["token"] = token
	ret["userinfo"] = userSchema
	return ret, nil
}

func ensureUser(res *models.WeChatLoginRes) *models.User {
	OpenID := res.OpenID
	user, err := daos.UserDAO.DoGetUserByOpenID(OpenID)
	if err != nil {
		user.Nickname = utils.GenRandomNickname()
		user = daos.UserDAO.DoCreateUser(OpenID)
	}
	return user
}

func setLoginTokenToRedis(key, token string) {
	context := context.Background()
	expire := time.Hour * 24 * 7
	err := core.RedisClient.SetEx(context, key, string(token), expire).Err()
	if err != nil {
		panic(err)
	}
}

func ensureWechatAuthToken() (string, error) {
	ctx := context.Background()
	token, err := core.RedisClient.Get(ctx, "wechat_access_token").Result()
	if err != nil {
		tokenRes, err := core.Client.GetAuthToken()
		if err != nil {
			return "", nil
		}
		expire := time.Second * time.Duration(tokenRes.ExpiresIn-100)
		err = core.RedisClient.SetEx(ctx, "wechat_access_token", tokenRes.AccessToken, expire).Err()
		if err != nil {
			return "", nil
		}
		token = tokenRes.AccessToken
	}
	return token, nil
}
