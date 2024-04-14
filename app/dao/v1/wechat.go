package dao_v1

import (
	"encoding/json"
	"fuge/app/core"
	models "fuge/app/models/v1"
	"net/url"
)

type wechatDAO struct {
	wechatLoginUrl string
	wechatPhoneUrl string
}

var WechatDAO *wechatDAO = &wechatDAO{
	wechatLoginUrl: "/sns/jscode2session?",
	wechatPhoneUrl: "/wxa/business/getuserphonenumber?",
}

func (s *wechatDAO) DoGetPhoneNumber(code string, authToken string) (*models.WechatPhoneRes, error) {
	data := &url.Values{
		"access_token": {authToken},
	}
	// data.Set("access_token", authToken)
	reqBody := map[string]string{
		"code": code,
	}
	bodyBytes, err := core.Client.Post(s.wechatPhoneUrl+data.Encode(), reqBody)
	if err != nil {
		return nil, err
	}
	body := &models.WechatPhoneRes{}
	if err := json.Unmarshal(bodyBytes, body); err != nil {
		return nil, err
	}
	return body, nil
}

func (s *wechatDAO) DoLoginWechat(lwi *models.LoginWeChatIn) *models.WeChatLoginRes {
	data := &url.Values{}
	conf := core.GetConf()
	data.Set("appid", conf.AppID)
	data.Set("secret", conf.AppSecret)
	data.Set("js_code", lwi.Code)
	data.Set("grant_type", "authorization_code")

	bodyBytes, err := core.Client.Get(s.wechatLoginUrl + data.Encode())
	if err != nil {
		panic(err)
	}
	body := &models.WeChatLoginRes{}
	if err := json.Unmarshal(bodyBytes, body); err != nil {
		panic(err)
	}
	// o2obk58PnZ-c40mACNtgtG3KYco4
	// SessionKey: MoUDMXCTjnOL3ZYTwSQzuA==
	if body.ErrorCode != 0 {
		panic(body.ErrMsg)
	}
	return body
}
