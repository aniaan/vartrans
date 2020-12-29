package trans

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"unicode/utf8"
)

type YouDaoResult struct {
	ErrorCode   string   `json:"errorCode"`
	Translation []string `json:"Translation"`
}

type YouDaoFetcher struct {
	appKey    string
	appSecret string
}

func (y *YouDaoFetcher) fetch(q string) ([]string, error) {
	api := "https://openapi.youdao.com/api"
	from := "zh-CHS"
	to := "en"
	signType := "v3"
	curtime := strconv.FormatInt(time.Now().Unix(), 10)
	salt := uuid.New().String()
	signStr := y.appKey + y.truncate(q) + salt + curtime + y.appSecret
	sign := y.encrypt(signStr)
	values := url.Values{
		"from":     {from},
		"to":       {to},
		"signType": {signType},
		"salt":     {salt},
		"sign":     {sign},
		"q":        {q},
		"appKey":   {y.appKey},
		"curtime":  {curtime},
	}
	resp, err := http.PostForm(api, values)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("网络请求失败")
	}

	body, err := ioutil.ReadAll(resp.Body)

	//暂时不管
	resp.Body.Close()

	if err != nil {
		return nil, errors.New("body读取失败")
	}
	var result YouDaoResult
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, errors.New("json加载失败")
	}

	if result.ErrorCode != "0" {
		return nil, errors.New("网络请求失败: errorCode=" + result.ErrorCode)
	}

	return result.Translation, nil
}

func (y *YouDaoFetcher) truncate(query string) string {
	r := []rune(query)
	size := utf8.RuneCountInString(query)

	if size < 20 {
		return query
	} else {
		return string(r[0:10]) + strconv.Itoa(size) + string(r[size-10:])
	}

}

func (y *YouDaoFetcher) encrypt(signStr string) string {
	value := sha256.Sum256([]byte(signStr))
	return hex.EncodeToString(value[:])

}

func NewYouDaoFetcher(appKey string, appSecret string) Fetcher {
	fetcher := YouDaoFetcher{
		appKey:    appKey,
		appSecret: appSecret,
	}
	return &fetcher
}
