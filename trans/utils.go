package trans

import (
	"strings"
)

var filterConfig map[string]int = map[string]int{
	"and": 1,
	"or":  1,
	"the": 1,
	"a":   1,
	"at":  1,
	"of":  1,
}

func preFilter(value string) []string {
	arr := strings.Split(value, " ")
	var newArr []string
	for _, item := range arr {
		if len(item) == 0 {
			continue
		}
		if _, ok := filterConfig[item]; ok {
			continue
		}
		newArr = append(newArr, strings.ToLower(item))
	}

	return newArr
}

// camelCase 驼峰命名法
func camelCase(value string) string {
	arr := preFilter(value)
	for i := 1; i < len(arr); i++ {
		arr[i] = strings.Title(arr[i])
	}
	return strings.Join(arr, "")
}

func bigCamelCase(value string) string {
	arr := preFilter(value)
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.Title(arr[i])
	}
	return strings.Join(arr, "")
}

func constant(value string) string {
	arr := preFilter(value)
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.ToUpper(arr[i])
	}
	return strings.Join(arr, "_")
}

func snakeCase(value string) string {
	arr := preFilter(value)
	return strings.Join(arr, "_")
}

func hyphen(value string) string {
	arr := preFilter(value)
	return strings.Join(arr, "-")
}
