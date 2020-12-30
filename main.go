package main

import (
	"errors"
	"github.com/BeanNan/vartrans/trans"
	"github.com/deanishe/awgo"
	"os"
	"strings"
)

var wf *aw.Workflow

func run() {
	fetcherName := os.Getenv("fetcher")
	fetcher, loadErr := loadFetcher(fetcherName)

	if loadErr != nil {
		wf.WarnEmpty("fetcher", loadErr.Error())
		wf.SendFeedback()
		return
	}

	method, q := os.Args[1], os.Args[2]

	results, err := runFetcher(q, method, fetcher)

	if err != nil {
		wf.WarnEmpty("fetcher", err.Error())
	} else if len(results) == 0 {
		wf.WarnEmpty("fetcher", "查询结果为空")
	} else {
		for _, result := range results {
			wf.NewItem(result.Title).
				Subtitle(result.SubTitle).
				Arg(result.Arg).
				Valid(true)
		}
	}
	wf.SendFeedback()
}

func loadFetcher(fetcherName string) (trans.Fetcher, error) {
	var fetcher trans.Fetcher
	var err error
	switch fetcherName {
	case trans.YOUDAO:
		fetcher, err = newYouDaoFetcher()
	default:
		fetcher, err = nil, errors.New("翻译API设置不正确")
	}

	return fetcher, err
}

func runFetcher(q string, method string, fetcher trans.Fetcher) ([]trans.Result, error) {
	translate := trans.Translate{
		Query:   q,
		Method:  method,
		Fetcher: fetcher,
	}
	results, err := translate.Execute()
	return results, err
}

func newYouDaoFetcher() (trans.Fetcher, error) {
	appKey := strings.TrimSpace(os.Getenv("appKey"))
	appSecret := strings.TrimSpace(os.Getenv("appSecret"))
	if appKey == "" || appSecret == "" {
		return nil, errors.New("youdao appKey和appSecret不能为空")
	}

	fetcher := trans.YouDaoFetcher{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	return &fetcher, nil
}

func main() {
	wf = aw.New()
	wf.Run(run)
}
