package star_map_ad_callback

import "testing"

// TestNewStarMapAdCallback 测试
func TestNewStarMapAdCallback(t *testing.T) {
	starMap := NewStarMapAdCallback()
	if starMap != nil {
		t.Log("实例化成功")
	} else {
		t.Error("实例化失败")
	}
}

// TestStarMapAdCallback_HttpGet 测试
func TestStarMapAdCallback_HttpGet(t *testing.T) {
	starMap := NewStarMapAdCallback()
	var res Response
	err := starMap.HttpGet(ApiHost, &res)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

// TestActivateCallback 测试
func TestActivateCallback(t *testing.T) {
	starMap := NewStarMapAdCallback()
	res, err := starMap.ActivateCallback("https://www.baidu.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

// TestRegisterCallback 测试
func TestRegisterCallback(t *testing.T) {
	starMap := NewStarMapAdCallback()
	res, err := starMap.RegisterCallback("https://www.baidu.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

// TestPayCallback 测试
func TestPayCallback(t *testing.T) {
	starMap := NewStarMapAdCallback()
	res, err := starMap.PayCallback("https://www.baidu.com", "{\"abc\":\"123\"}")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
