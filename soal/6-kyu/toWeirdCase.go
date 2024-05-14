// Question :

// Write a function that accepts a string, and returns the same string with all even indexed characters in each word upper cased, and all odd indexed characters in each word lower cased. The indexing just explained is zero based, so the zero-ith index is even, therefore that character should be upper cased and you need to start over for each word.
// The passed in string will only consist of alphabetical characters and spaces(' '). Spaces will only be present if there are multiple words. Words will be separated by a single space(' ').
// Examples:
// "String" => "StRiNg"
// "Weird string case" => "WeIrD StRiNg CaSe"

// Answer :
package soal

import "strings"

func ToWeirdCase(str string) string {

	arr := strings.Split(str, " ")
	newArr := []string{}

	for i := 0; i < len(arr); i++{
		newStr := ""
		for j := 0; j < len(arr[i]); j++{
			if j % 2 == 0 {
				cap := strings.ToUpper(string(arr[i][j]))
				newStr += cap
			} else {
				normal := strings.ToLower(string(arr[i][j]))
				newStr += normal
			}
		}
		newArr = append(newArr, newStr)
	}

	return strings.Join(newArr, " ")

}