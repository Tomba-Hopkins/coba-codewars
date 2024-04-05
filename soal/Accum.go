// Question :

// This time no story, no theory. The examples below show you how to write function accum:
// Examples:
// accum("abcd") -> "A-Bb-Ccc-Dddd"
// accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
// accum("cwAt") -> "C-Ww-Aaa-Tttt"
// The parameter of accum is a string which includes only letters from a..z and A..Z.

// Answer :
package soal

import "strings"

func Accum(s string) (res string) {

	for i := 0; i < len(s); i++{
		res += strings.ToUpper(string(s[i]))
		for j := 0; j < i; j++{
			res += strings.ToLower(string(s[i]))
		}
		res += "-"
	}

	res = res[:len(res) - 1]

	return


}