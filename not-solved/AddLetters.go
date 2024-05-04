// Question :

// Your task is to add up letters to one letter.

// The function will be given a slice of letters (runes), each one being a letter to add. Return a rune.

// Notes:
// Letters will always be lowercase.
// Letters can overflow (see second to last example of the description)
// If no letters are given, the function should return 'z'
// Examples:
// AddLetters([]rune{'a', 'b', 'c'}) = 'f'
// AddLetters([]rune{'a', 'b'}) = 'c'
// AddLetters([]rune{'z'}) = 'z'
// AddLetters([]rune{'z', 'a'}) = 'a'
// AddLetters([]rune{'y', 'c', 'b'}) = 'd' // notice the letters overflowing
// AddLetters([]rune{}) = 'z'

// Answer :
package soal

import "fmt"

func AddLetters(letters []rune) rune {

	kodingkanDuluLe := "zabcdefgihjklmnopqrstuvwxy"
	abjad := map[rune]int{}

	for i, k := range kodingkanDuluLe{
		abjad[k] = i
	}


	dahDikodingkanNihLe := map[int]rune{}

	for j := 0; j < len(kodingkanDuluLe); j++{
		dahDikodingkanNihLe[j] = rune(kodingkanDuluLe[j])
	}


	total := 0
	
	for l := 0; l < len(letters); l++{
		for r := 0; r < abjad[letters[l]]; r++{
			total++
		}
	}

	fmt.Println(string(dahDikodingkanNihLe[total]))


	return dahDikodingkanNihLe[total]
	

}