package star_map_ad_callback

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

// ApiHost 广告回调地址
const (
	ApiHost = "https://ad.oceanengine.com/track/activate/"
)

// event_type=0表示激活，event_type=1表示注册；event_type=2表示付费，
const (
	EventTypeActivate = 0
	EventTypeRegister = 1
	EventTypePay      = 2
)

// Response 返回数据
type Response struct {
	Code int    `json:"code,omitempty"`
	Ret  int    `json:"ret,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

// StarMapAdCallback 广告回调
type StarMapAdCallback struct {
}

// NewStarMapAdCallback 实例化
func NewStarMapAdCallback() *StarMapAdCallback {
	return &StarMapAdCallback{}
}

// ActivateCallback 激活
func (s *StarMapAdCallback) ActivateCallback(callback string) (res Response, err error) {
	url := fmt.Sprintf("%s?event_type=%d&callback=%s", ApiHost, EventTypeActivate, callback)
	err = s.HttpGet(url, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// RegisterCallback 注册
func (s *StarMapAdCallback) RegisterCallback(callback string) (res Response, err error) {
	url := fmt.Sprintf("%s?event_type=%d&callback=%s", ApiHost, EventTypeRegister, callback)
	err = s.HttpGet(url, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// PayCallback 付费
func (s *StarMapAdCallback) PayCallback(callback string, props string) (res Response, err error) {
	url := fmt.Sprintf("%s?event_type=%d&callback=%s&props=%s", ApiHost, EventTypePay, callback, url2.QueryEscape(props))
	err = s.HttpGet(url, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// HttpGet http get请求
func (s *StarMapAdCallback) HttpGet(url string, res interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, res)
	if err != nil {
		return err
	}
	return nil
}
