package models_v1

import (
// "github.com/gin-gonic/gin/binding"
// "github.com/go-playground/validator/v10"
)

type WechatResBase struct {
	ErrorCode int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

type LoginWeChatIn struct {
	Code string `json:"code" binding:"required" type:"string" format:"string" description:"wechat code" nullable:"false"`
}

type WeChatLoginRes struct {
	WechatResBase
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

type AuthTokenSchema struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// type PhoneInfo struct {
// 	PhoneNumber     string `json:"phoneNumber"`
// 	PurePhoneNumber string `json:"purePhoneNumber"`
// 	CountryCode     string `json:"countryCode"`
// 	Watermark       struct {
// 		AppID     string `json:"appid"`
// 		Timestamp int    `json:"timestamp"`
// 	} `json:"watermark"`
// }

type WechatPhoneRes struct {
	WechatResBase
	PhoneInfo struct {
		PhoneNumber     string `json:"phoneNumber"`
		PurePhoneNumber string `json:"purePhoneNumber"`
		CountryCode     string `json:"countryCode"`
		Watermark       struct {
			AppID     string `json:"appid"`
			Timestamp int    `json:"timestamp"`
		} `json:"watermark"`
	} `json:"phone_info"`
}
