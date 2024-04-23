// Question :

// Sort the Vowels!
// In this kata, we want to sort the vowels in a special format.
// Task
// Write a function which takes a input string s and return a string in the following way:
//                   C|                          R|
//                   |O                          n|
//                   D|                          d|
//    "CODEWARS" =>  |E       "Rnd Te5T"  =>      |
//                   W|                          T|
//                   |A                          |e
//                   R|                          5|
//                   S|                          T|
// Notes:
// List all the Vowels on the right side of |
// List every character except Vowels on the left side of |
// for the purpose of this kata, the vowels are : a e i o u
// Return every character in its original case
// Each line is seperated with \n
// Invalid input ( undefined / null / integer ) should return an empty string

// Answer :
package soal

import "strings"

func SortVowels(s string) (res string) {

	abjad := map[string]bool{
		"a": true,
		"i": true,
		"u": true,
		"e": true,
		"o": true,
	}

	for _, r := range s {
		rString := string(r)

		if rString == " " {
			res += " |\n"
		} else if abjad[strings.ToLower(rString)] {
			res +=  "|" + rString + "\n"
		} else {
			res += rString + "|\n"
		}
	}

	if len(res) > 20 {
		res = res[:len(res) - 1]
	}
	
	return 
}