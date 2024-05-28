// Question :

// In this Kata, you will be given a string that may have mixed uppercase and lowercase letters and your task is to convert that string to either lowercase only or uppercase only based on:
// make as few changes as possible.
// if the string contains equal number of uppercase and lowercase letters, convert the string to lowercase.
// For example:
// solve("coDe") = "code". Lowercase characters > uppercase. Change only the "D" to lowercase.
// solve("CODe") = "CODE". Uppercase characters > lowecase. Change only the "e" to uppercase.
// solve("coDE") = "code". Upper == lowercase. Change all to lowercase.

// Answer :
package soal

import "strings"

func FixStringCase(str string) string {
	text := "abcdefghijklmnopqrstuvwxyz"
	abjad := map[string]string{}

	for _, t := range text {
		abjad[string(t)] = string(t)
	}

	up, low := 0, 0

	for _, r := range str {
		if string(r) == abjad[string(r)] {
			low++
		} else {
			up++
		}
	}

	if up > low{
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)

}