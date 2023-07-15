package utils

import "fmt"

func JoinString(s1 string, s2 string) string {
	return fmt.Sprintf("%s: %s", s1, s2)
}
