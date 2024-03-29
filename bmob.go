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

	if isPay, has := cc.Get(projectName); has {
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

	req := httplib.Get(`https://api2.bmob.cn/1/classes/App?where={"name":"` + projectName + `"}`)
	req.Header("X-Bmob-Application-Id", "a7517b83302c544bac2b9cbd7a7d7fa1")
	req.Header("X-Bmob-REST-API-Key", "43d6eb0aa1236955a1c9fa2a575e7ad6")
	req.Header("Content-Type", "application/json")
	err := req.ToJSON(&r)
	if err != nil {
		cc.Set(projectName, true, 20*time.Minute)
		return true
	}
	if len(r.Results) == 0 {
		cc.Set(projectName, true, 20*time.Minute)
		return true
	}
	if r.Results[0].IsPay {
		cc.Set(projectName, true, 2*time.Hour)
		return true
	}
	if time.Now().Format("2006-01-02 15:04:05") > r.Results[0].ExpireDate.Iso {
		cc.Set(projectName, false, 20*time.Minute)
		return false
	}
	cc.Set(projectName, true, 20*time.Minute)
	return true
}
