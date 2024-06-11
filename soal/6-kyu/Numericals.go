// Question :

// You are given an input string.
// For each symbol in the string if it's the first character occurrence, replace it with a '1', else replace it with the amount of times you've already seen it.
// Examples:
// input   =  "Hello, World!"
// result  =  "1112111121311"
// input   =  "aaaaaaaaaaaa"
// result  =  "123456789101112"
// There might be some non-ascii characters in the string.
// Take note of performance

// Answer :
package soal

import (
	"strconv"
	"strings"
)

func Numericals(s string) string {

	var res strings.Builder

	abjad := map[rune]int{}

	for _, r := range s {

		abjad[r]++
		res.WriteString(strconv.Itoa(abjad[r]))
	}

	return res.String()

}

// Before
// func Numericals(s string) (res string) {
// 	abjad := map[rune]int{}
// 	for _, r := range s {
// 		abjad[r]++
// 		res += strconv.Itoa(abjad[r])
// 	}
// 	return
// }