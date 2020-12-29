package trans

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	config := YouDaoAuthConfig{AppKey: "", AppSecret: ""}
	trans := Translate{AuthConfig: config, Query: "你好", Method: "zh"}

	result, err := trans.Execute()

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Print(err)
	}
}
