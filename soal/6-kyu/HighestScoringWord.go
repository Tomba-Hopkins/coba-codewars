// Question :

// Given a string of words, you need to find the highest scoring word.
// Each letter of a word scores points according to its position in the alphabet: a = 1, b = 2, c = 3 etc.
// For example, the score of abad is 8 (1 + 2 + 1 + 4).
// You need to return the highest scoring word as a string.
// If two words score the same, return the word that appears earliest in the original string.
// All letters will be lowercase and all inputs will be valid.

// Answer :
package soal

import "strings"

func High(s string) (res string) {

	arr := strings.Split(s, " ")
	str := "abcdefghijklmnopqrstuvwxyz"

	abjad := map[string]int{}
	
	for index, sr := range str{
		abjad[string(sr)] = index + 1
	}

	resArr := []int{}

	for i := 0; i < len(arr); i++{
		temp := 0
		kata := arr[i]

		for j := 0; j < len(arr[i]); j++{
			char := kata[j]
			temp += abjad[string(char)]
		}
		resArr = append(resArr, temp)
		temp = 0
	}

	
	high := 0

	for k := len(resArr) - 1; k >= 0; k--{
		if resArr[k] >= high{
			high = resArr[k]
			res = arr[k]
		}
	} 


	return 
	
}