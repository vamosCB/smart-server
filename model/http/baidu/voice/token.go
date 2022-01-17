package voice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	// 百度智能云平台获取的 API Key
	ClientID = "zz4Q8E4V5tlWPcNyPMLfPAKe"
	// 百度智能云平台获取的 Secret Key
	ClientSecret = "TMPFeuXNRApEX6R5KrLRtbFevN71jlNd"
	// 获取token的地址
	Host = "https://aip.baidubce.com/oauth/2.0/token"
)

// Token 结构体
type Token struct {
	Err              string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ExpiresIn        int    `json:"expires_in"` //Access Token的有效期(秒为单位，有效期30天)；
	AccessToken      string `json:"access_token"`
	SessionSecret    string `json:"session_secret"`
}

// GetToken 获取鉴权
func GetToken() (*Token, error) {
	result := &Token{}
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", ClientID)
	data.Set("client_secret", ClientSecret)
	httpResp, err := http.PostForm(Host, data)
	if err != nil {
		return result, fmt.Errorf("Request Fail, err=%+v", err.Error())
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return result, fmt.Errorf("Read Resp Body Fail, err=%+v", err.Error())
	}
	if err := json.Unmarshal(body, result); err != nil {
		return result, fmt.Errorf("Request Fail, err=%+v", err.Error())
	}
	return result, nil

}
