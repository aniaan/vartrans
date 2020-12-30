package trans

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockFetcher struct {
	mock.Mock
}

func (m *MockFetcher) fetch(q string) ([]string, error) {
	args := m.Called(q)
	return args.Get(0).([]string), args.Error(1)
}

func TestTranslate(t *testing.T) {
	tests := []struct {
		q            string
		method       string
		fetchResults []string
		fetchErr     error
		transResults []Result
		transErr     error
		hasErr       bool
		err          error
	}{
		{
			q:            "你好，世界",
			method:       CL,
			fetchResults: []string{"hello world"},
			transResults: []Result{{Title: "HELLO_WORLD", SubTitle: "标准翻译 => hello world", Arg: "HELLO_WORLD"}},
			hasErr:       false,
		},
		{
			q:            "你和我",
			method:       XH,
			fetchResults: []string{"you and me"},
			transResults: []Result{{Title: "you_me", SubTitle: "标准翻译 => you and me", Arg: "you_me"}},
			hasErr:       false,
		},
		{
			q:            "你和我",
			method:       XT,
			fetchResults: []string{"you and me"},
			transResults: []Result{{Title: "youMe", SubTitle: "标准翻译 => you and me", Arg: "youMe"}},
			hasErr:       false,
		},
		{
			q:            "你和我",
			method:       DT,
			fetchResults: []string{"you and me"},
			transResults: []Result{{Title: "YouMe", SubTitle: "标准翻译 => you and me", Arg: "YouMe"}},
			hasErr:       false,
		},
		{
			q:            "你和我",
			method:       ZH,
			fetchResults: []string{"you and me"},
			transResults: []Result{{Title: "you-me", SubTitle: "标准翻译 => you and me", Arg: "you-me"}},
			hasErr:       false,
		},
		{
			q:            "你和我",
			method:       ZH,
			fetchResults: nil,
			fetchErr:     errors.New("网络异常"),
			transResults: nil,
			transErr:     errors.New("网络异常"),
			hasErr:       true,
			err:          errors.New("网络异常"),
		},
		{
			q:      "你和我",
			method: "none",
			hasErr: true,
			err:    errors.New("invalid method"),
		},
	}

	for index, test := range tests {
		mockFetcher := new(MockFetcher)
		mockFetcher.On("fetch", test.q).Return(test.fetchResults, test.fetchErr)

		translate := Translate{
			Query:   test.q,
			Method:  test.method,
			Fetcher: mockFetcher,
		}
		results, err := translate.Execute()

		if test.hasErr {
			assert.EqualError(t, err, test.err.Error(), "equal error row %d", index)
		} else {
			assert.NoError(t, err, "Should not be error row %d", index)
		}

		assert.Equal(t, test.transResults, results)
	}

}
