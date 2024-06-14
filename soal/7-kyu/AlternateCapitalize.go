// Question :

// Given a string, capitalize the letters that occupy even indexes and odd indexes separately, and return as shown below. Index 0 will be considered even.
// For example, capitalize("abcdef") = ['AbCdEf', 'aBcDeF']. See test cases for more examples.
// The input will be a lowercase string with no spaces.
// Good luck!

// Answer :
package soal

import "strings"

func AlternateCapitalize(st string) (res []string) {

	a, b := "", ""

	for i := 0; i < len(st); i++{
		if i % 2 == 0 {
			a += strings.ToUpper(string(st[i]))
		} else {
			a += string(st[i])
		}
	}

	for i := 0; i < len(st); i++{
		if i % 2 != 0 {
			b += strings.ToUpper(string(st[i]))
		} else {
			b += string(st[i])
		}
	}

	res = append(res, a, b)

	return

}