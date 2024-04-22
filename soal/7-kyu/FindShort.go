// Question :
// Simple, given a string of words, return the length of the shortest word(s).
// String will never be empty and you do not need to account for different data types.

// Answer :
package soal

import "strings"

func FindShort(s string) int {
	

	arr := strings.Split(s, " ")

	result := len(arr[0])

	for i := 0; i < len(arr); i++{
		if len(arr[i]) < result {
			result = len(arr[i])
		}
	}

	return result
}
