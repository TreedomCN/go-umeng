package umeng_test

import (
	"github.com/netroby/umeng-go-sdk"
	"testing"
)

var data *umeng.Data

func init() {
	data = umeng.NewData(umeng.AppAndroid)
}

func TestStatus(t *testing.T) {
	resp := data.Status()
	if len(resp.Code) > 0 {
		t.Log("data.Status 测试通过")
	} else {
		t.Error("data.Status 测试不通过")
	}

}

func TestCancel(t *testing.T) {
	resp := data.Cancel()
	if len(resp.Code) > 0 {
		t.Log("data.Cancel 测试通过")
	} else {
		t.Error("data.Cancel 测试不通过")
	}
}

func TestUpload(t *testing.T) {
	resp := data.Upload()
	if len(resp.Code) > 0 {
		t.Log("data.Upload 测试通过")
	} else {
		t.Error("data.Upload 测试不通过")
	}
}

func TestPush(t *testing.T) {
	body := umeng.AndroidBody{}
	body.DisplayType = "message"
	body.Custom = "test"
	body.Title = "1212"

	extra := make(map[string]string, 0)
	extra["key1"] = "1212x"
	extra["key2"] = "12123323"
	resp := data.Push(body, nil, nil, extra)
	if len(resp.Code) > 0 {
		t.Log("data.Push 测试通过")
	} else {
		t.Error("data.Push 测试不通过")
	}
}
