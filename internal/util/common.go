package util

import (
	"fmt"
	"strings"
)

func GetOrderByFromString(s string) string {
	if s == "" {
		return s
	}
	var result []string
	arrOrderBy := strings.Split(s, ",")
	for _, item := range arrOrderBy {
		order := item[0]
		str := item[1:]

		if string(order) == "+" {
			result = append(result, fmt.Sprintf("%s ASC", str))
		} else if string(order) == "-" {
			result = append(result, fmt.Sprintf("%s DESC", str))
		}
	}

	return strings.Join(result, ",")
}
