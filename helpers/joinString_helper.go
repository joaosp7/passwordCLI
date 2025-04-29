package helpers

import "strings"

func PassowrdsToString(passwords []string) string{
	var builder strings.Builder
	for _, str := range passwords{
		builder.WriteString(str)
		builder.WriteString("\n")
	}
	result := builder.String()

	return result
}