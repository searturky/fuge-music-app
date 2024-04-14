package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	models "fuge/app/models/v1"
	"io"
	"net/http"
	"net/url"
)

var rawClient = &http.Client{}
var validStatus = initValidHttpStatus()

type WechatClient struct {
	RawClient   *http.Client
	Host        string
	Scheme      string
	ContentType string
}

var Client *WechatClient = &WechatClient{
	RawClient:   rawClient,
	Host:        "api.weixin.qq.com",
	Scheme:      "https",
	ContentType: "application/json",
}

func (s *WechatClient) Get(url string) ([]byte, error) {
	fullUrl := s.Scheme + "://" + s.Host + url
	resp, err := s.RawClient.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := handleResponse(resp)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func (s *WechatClient) Post(url string, body map[string]string) ([]byte, error) {
	fullUrl := s.Scheme + "://" + s.Host + url
	bodyJson, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	b := bytes.NewBuffer(bodyJson)
	resp, err := s.RawClient.Post(fullUrl, s.ContentType, b)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func (s *WechatClient) PostForm(url string, data map[string][]string) (*http.Response, error) {
	return s.RawClient.PostForm(url, data)
}

func (s *WechatClient) Put(targetUrl string, body io.Reader) (*http.Response, error) {
	url := url.URL{
		Scheme: s.Scheme,
		Host:   s.Host,
		Path:   targetUrl,
	}
	req := &http.Request{
		Method: http.MethodPut,
		URL:    &url,
		Header: map[string][]string{
			"Content-Type": {s.ContentType},
		},
		Body: io.NopCloser(body),
	}
	return s.RawClient.Do(req)
}

func (s *WechatClient) Delete(targetUrl string) (*http.Response, error) {
	url := url.URL{
		Scheme: s.Scheme,
		Host:   s.Host,
		Path:   targetUrl,
	}
	req := &http.Request{
		Method: http.MethodDelete,
		URL:    &url,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}
	return s.RawClient.Do(req)
}

func (s *WechatClient) GetAuthToken() (*models.AuthTokenSchema, error) {
	targetUrl := url.URL{
		Path: "/cgi-bin/token",
	}
	coreConf := GetConf()
	query := &url.Values{}
	query.Set("grant_type", "client_credential")
	query.Set("appid", coreConf.AppID)
	query.Set("secret", coreConf.AppSecret)
	targetUrl.RawQuery = query.Encode()
	bodyBytes, err := s.Get(targetUrl.String())
	if err != nil {
		return nil, err
	}
	body := &models.AuthTokenSchema{}
	if err := json.Unmarshal(bodyBytes, body); err != nil {
		return nil, err
	}
	return body, nil
}

func initValidHttpStatus() []int {
	v := []int{}
	for i := 200; i < 300; i++ {
		v = append(v, i)
	}
	return v
}

func isValidHttpStatus(resp *http.Response) bool {
	for _, status := range validStatus {
		if resp.StatusCode == status {
			return true
		}
	}
	return false
}

func handleResponse(resp *http.Response) ([]byte, error) {
	if !isValidHttpStatus(resp) {
		return nil, errors.New("invalid http status")
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyBase := &models.WechatResBase{}
	if err := json.Unmarshal(resBody, bodyBase); err != nil {
		return nil, err
	}
	if bodyBase.ErrorCode != 0 {
		return nil, errors.New(bodyBase.ErrMsg)
	}
	return resBody, nil
}
