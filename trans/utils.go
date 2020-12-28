package trans

import "strings"


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
		if _, ok := filterConfig[item]; ok {
			continue
		}

		newArr = append(newArr, item)
	}

	return []string{}
}

func camelCase(value string) string {
	return ""
}

func bigCamelCase(value string) string {
	return ""
}

func constant(value string) string {
	return ""
}

func snakeCase(value string) string {
	return ""
}

func hyphen(value string) string {
	return ""
}
