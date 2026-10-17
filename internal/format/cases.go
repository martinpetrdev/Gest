package format

import "strings"

func PascalCase(strs []string) string {
	var out strings.Builder
	for _, str := range strs {
		out.WriteString(Capitalize(str))
	}

	return out.String()
}

func Capitalize(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}
