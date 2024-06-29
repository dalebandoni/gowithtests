package iteration

import "strings"

func Repeat(character string, repititions int) string {
	var repeated string
	for i := 0; i < repititions; i++ {
		repeated += character
	}
	return repeated
}

func GetIndex(character, substring string) (answer int) {
	answer = strings.Index(character, substring)
	return
}
