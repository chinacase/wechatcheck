package wechatcheck

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	//StatusUnknown 未知
	StatusUnknown = 2
	//StatusPass 通过
	StatusPass = 1
	//StatusNoPass 不通过
	StatusNoPass = 0
)

//Result 返回信息
type Result struct {
	Status int    //状态 1 通过 0 不通过 2 未知
	Msg    string //提示
}

//CheckURL 检测地址是否通过微信
func CheckURL(urlString string) (result Result) {
	var getURL string
	if !strings.HasPrefix(urlString, "http://") && !strings.HasPrefix(urlString, "https://") {
		urlString = "http://" + urlString
	}
	getURL = "https://mp.weixinbridge.com/mp/wapredirect?url=" + urlString
	client := http.Client{
		Timeout: time.Second * 2,
	}
	result = Result{}
	req, _ := http.NewRequest("GET", getURL, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4406.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		result.Status = StatusUnknown
		result.Msg = "Unknown"
		return
	}
	defer resp.Body.Close()
	locltionURL := resp.Request.URL.String()

	localtionP, err1 := url.Parse(locltionURL)

	if err1 != nil {
		result.Status = StatusUnknown
		result.Msg = "Unknown"
		return
	}

	urlP, err2 := url.Parse(urlString)
	if err2 != nil {
		result.Status = StatusUnknown
		result.Msg = "Unknown"
		return
	}

	if strings.Compare(strings.ToLower(urlP.Host), strings.ToLower(localtionP.Host)) == 0 {
		result.Status = StatusPass
		result.Msg = "Pass"
		return
	}
	result.Status = StatusNoPass
	result.Msg = "NoPass"
	return
}
