package bmob

import (
	"time"
	"github.com/astaxie/beego/httplib"
)

func IsPay(projectName string)bool{
	type Result struct {
		Results []struct{
			IsPay bool `json:"isPay"`
			ExpireDate struct{
				Iso string `json:"iso"`
			} `json:"expireDate"`
		} `json:"results"`
	}
	var r Result

	req:=httplib.Get(`https://api2.bmob.cn/1/classes/project?where={"name":"`+projectName+`"}`)
	req.Header("X-Bmob-Application-Id","9cc2fc47276419893dc4361352b3cef0")
	req.Header("X-Bmob-REST-API-Key","405c5d5f1296375287d908a964d1504c")
	req.Header("Content-Type","application/json")
	err:=req.ToJSON(&r)
	if err!=nil{
		return true
	}
	if len(r.Results)==0{
		return true
	}
	if r.Results[0].IsPay{
		return true
	}
	if time.Now().Format("2006-01-02 15:04:05")>r.Results[0].ExpireDate.Iso{
		return false
	}
	return true
}