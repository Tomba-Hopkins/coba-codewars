// Question :

// The goal of this exercise is to convert a string to a new string where each character in the new string is "(" if that character appears only once in the original string, or ")" if that character appears more than once in the original string. Ignore capitalization when determining if a character is a duplicate.
// Examples
// "din"      =>  "((("
// "recede"   =>  "()()()"
// "Success"  =>  ")())())"
// "(( @"     =>  "))(("
// Notes
// Assertion messages may be unclear about what they display in some languages. If you read "...It Should encode XXX", the "XXX" is the expected result, not the input!

// Answer :
package soal

import "strings"

func DuplicateEncode(word string) (res string) {

	abjad := map[string]int{}

	for _, w := range word {
		new := strings.ToLower(string(w))
		abjad[new]++
	}

	for _, r := range word {
		newR := strings.ToLower(string(r))
		if abjad[newR] > 1 {
			res += ")"
		} else {
			res += "("
		}
	}

	return
}

// func DuplicateEncode(word string) (res string) {

// 	word = strings.ToLower(word)

// 	abjad := map[string]int{}

// 	for _, w := range word {
// 		abjad[string(w)]++
// 	}

// 	for _, r := range word {
// 		if abjad[string(r)] > 1 {
// 			res += ")"
// 		} else {
// 			res += "("
// 		}
// 	}

// 	return
// }