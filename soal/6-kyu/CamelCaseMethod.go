// Question :

// Write a method (or function, depending on the language) that converts a string to camelCase, that is, all words must have their first letter capitalized and spaces must be removed.
// Examples (input --> output):
// "hello case" --> "HelloCase"
// "camel case word" --> "CamelCaseWord"
// Don't forget to rate this kata! Thanks :)

// Answer :
package soal

import "strings"

func CamelCase(s string) (res string) {
	arr := strings.Split(s, " ")

	for i := 0; i < len(arr); i++{
		for j := 0; j < len(arr[i]); j++{
			if j == 0 {
				res += strings.ToUpper(string(arr[i][j]))
			} else {
				res += string(arr[i][j])
			}
		}
	}

	return
}
