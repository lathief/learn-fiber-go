package utils

import "fmt"

func JoinStringWithColon(s1 string, s2 string) string {
	if s1 == "" {
		return s2
	}
	return fmt.Sprintf("%s: %s", s1, s2)
}
func JoinStringWithSpace(s1 string, s2 string) string {
	if s1 == "" {
		return s2
	}
	return fmt.Sprintf("%s %s", s1, s2)
}
func JoinStringWithComma(s1 string, s2 string) string {
	if s1 == "" {
		return s2
	}
	return fmt.Sprintf("%s, %s", s1, s2)
}
func JoinStringWithStop(s1 string, s2 string) string {
	if s1 == "" {
		return s2
	}
	return fmt.Sprintf("%s. %s", s1, s2)
}
