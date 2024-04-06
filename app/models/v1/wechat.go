package models_v1

import (
// "github.com/gin-gonic/gin/binding"
// "github.com/go-playground/validator/v10"
)

type LoginWeChatIn struct {
	Code string `json:"code" binding:"required" type:"string" format:"string" description:"wechat code" nullable:"false"`
}

type LoginWeChatRes struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrorCode  int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
