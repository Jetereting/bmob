package bmob

import (
	"github.com/astaxie/beego/httplib"
	"github.com/maemual/go-cache"
	"time"
)

var cc *cache.Cache

func init() {
	cc = cache.New(10*time.Hour, 10*time.Hour)
}

func IsPay(projectName string) bool {

	if isPay, has := cc.Get("IsPay"); has {
		return isPay.(bool)
	}

	type Result struct {
		Results []struct {
			IsPay      bool `json:"isPay"`
			ExpireDate struct {
				Iso string `json:"iso"`
			} `json:"expireDate"`
		} `json:"results"`
	}
	var r Result

	req := httplib.Get(`https://api2.bmob.cn/1/classes/project?where={"name":"` + projectName + `"}`)
	req.Header("X-Bmob-Application-Id", "9cc2fc47276419893dc4361352b3cef0")
	req.Header("X-Bmob-REST-API-Key", "405c5d5f1296375287d908a964d1504c")
	req.Header("Content-Type", "application/json")
	err := req.ToJSON(&r)
	if err != nil {
		cc.Set("IsPay", true, 20*time.Minute)
		return true
	}
	if len(r.Results) == 0 {
		cc.Set("IsPay", true, 20*time.Minute)
		return true
	}
	if r.Results[0].IsPay {
		cc.Set("IsPay", true, 2*time.Hour)
		return true
	}
	if time.Now().Format("2006-01-02 15:04:05") > r.Results[0].ExpireDate.Iso {
		cc.Set("IsPay", false, 20*time.Minute)
		return false
	}
	cc.Set("IsPay", true, 20*time.Minute)
	return true
}
