package word

import (
	"strings"
	"unicode"
)

//全部转为大/小写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

//下划线单词转换为大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, data := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(rune(s[0])))
			continue
		}
		if unicode.IsUpper(data) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(rune(data)))
	}
	return string(output)
}
