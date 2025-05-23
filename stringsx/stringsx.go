package stringsx

import "strings"

func Split(str1 string, str2 string) []string {
	strs := strings.Split(str1, str2)
	if len(strs) == 0 || len(strs) == 1 && strs[0] == "" {
		return []string{}
	}
	return strs
}

func StripSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}
