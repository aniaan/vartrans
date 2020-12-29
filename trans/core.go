package trans

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

// Translate translate query
type Translate struct {
	AuthConfig YouDaoAuthConfig
	Query      string
	Method     string
}

func (t *Translate) fetch() ([]string, error) {
	api := "https://openapi.youdao.com/api"
	from := "zh-CHS"
	to := "en"
	signType := "v3"
	curtime := strconv.FormatInt(time.Now().Unix(), 10)
	salt := uuid.New().String()
	signStr := t.AuthConfig.AppKey + t.truncate(t.Query) + salt + curtime + t.AuthConfig.AppSecret
	sign := t.encrypt(signStr)
	values := url.Values{
		"from":     {from},
		"to":       {to},
		"signType": {signType},
		"salt":     {salt},
		"sign":     {sign},
		"q":        {t.Query},
		"appKey":   {t.AuthConfig.AppKey},
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

func (t *Translate) Execute() ([]Result, error) {
	switch t.Method {
	case CL:
		return t.run(constant)
	case XT:
		return t.run(snakeCase)
	case DT:
		return t.run(bigCamelCase)
	case XH:
		return t.run(snakeCase)
	case ZH:
		return t.run(hyphen)
	default:
		return []Result{}, errors.New("method")
	}

}

func (t *Translate) run(f func(value string) string) ([]Result, error) {
	fetchResult, err := t.fetch()

	if err != nil {
		return nil, err
	}
	var data []Result
	for _, item := range fetchResult {
		value := f(item)
		data = append(data, Result{
			Title:    value,
			SubTitle: "标准翻译 => " + item,
			Arg:      value,
		})
	}

	return data, nil
}

func (t *Translate) truncate(query string) string {
	r := []rune(query)
	size := utf8.RuneCountInString(query)

	if size < 20 {
		return query
	} else {
		return string(r[0:10]) + strconv.Itoa(size) + string(r[size-10:])
	}

}

func (t Translate) encrypt(signStr string) string {
	value := sha256.Sum256([]byte(signStr))
	return hex.EncodeToString(value[:])

}
