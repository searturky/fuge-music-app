package dao_v1

import (
	"encoding/json"
	"fuge/app/core"
	models "fuge/app/models/v1"
	"io"
	"net/url"
)

type wechatDAO struct {
}

var WechatDAO *wechatDAO = &wechatDAO{}

func (s *wechatDAO) DoLoginWechat(lwi *models.LoginWeChatIn) *models.LoginWeChatRes {
	data := &url.Values{}
	conf := core.GetConf()
	data.Set("appid", conf.AppID)
	data.Set("secret", conf.AppSecret)
	data.Set("js_code", lwi.Code)
	data.Set("grant_type", "authorization_code")

	resp, err := core.Client.Get("https://api.weixin.qq.com/sns/jscode2session?" + data.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyJson := &models.LoginWeChatRes{}
	if err := json.Unmarshal(body, bodyJson); err != nil {
		panic(err)
	}
	// o2obk58PnZ-c40mACNtgtG3KYco4
	// SessionKey: MoUDMXCTjnOL3ZYTwSQzuA==
	if bodyJson.ErrorCode != 0 {
		panic(bodyJson.ErrMsg)
	}
	println(bodyJson)
	return bodyJson
}
