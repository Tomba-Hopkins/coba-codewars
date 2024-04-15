// Question :

// Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
// The output should be two capital letters with a dot separating them.
// It should look like this:
// Sam Harris => S.H
// patrick feeney => P.F

// Answer :
package soal

import "strings"

func AbbrevName(name string) (res string) {

	box := strings.Split(name, " ")

	for i := 0; i < len(box); i++{
		cap := strings.ToUpper(string(box[i]))
		res += string(cap[0]) + "."
	}

	res = res[:len(res) - 1]

	return 
}