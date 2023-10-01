package general

import (
	"strconv"
)

func ValidateParam(id string) (uint, error) {
	var intId int
	var err error
	intId, err = strconv.Atoi(id)
	return uint(intId), err
}

func IsInArray(word string, words []string) bool {
	var result bool = false
	for i := range words {
		if words[i] == word {
			result = true
			break
		}
	}
	return result
}
