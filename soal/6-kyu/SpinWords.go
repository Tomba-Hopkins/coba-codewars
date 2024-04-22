// Question :

// Write a function that takes in a string of one or more words, and returns the same string, but with all words that have five or more letters reversed (Just like the name of this Kata). Strings passed in will consist of only letters and spaces. Spaces will be included only when more than one word is present.
// Examples:
// "Hey fellow warriors"  --> "Hey wollef sroirraw"
// "This is a test        --> "This is a test"
// "This is another test" --> "This is rehtona test"

// Answer :
package soal

import "strings"

func SpinWords(str string) (res string) {
	kotak := strings.Split(str, " ")
	
	for i := 0; i < len(kotak); i++{
		if len(kotak[i]) > 4 {
			for j := len(kotak[i]) - 1; j >= 0; j--{
				res += string(kotak[i][j])
			} 
			res += " "
		} else {
			res += kotak[i] + " "
		}
	}
	
	res = res[:len(res) - 1]

	return 
}