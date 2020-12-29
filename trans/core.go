package trans

import (
	"errors"
)

type Result struct {
	Title string
	SubTitle string
	Arg string
}

// Translate translate query
type Translate struct {
	Query   string
	Method  string
	fetcher Fetcher
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
	fetchResult, err := t.fetcher.fetch(t.Query)

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
