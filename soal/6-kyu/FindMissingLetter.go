// Question :

// Find the missing letter
// Write a method that takes an array of consecutive (increasing) letters as input and that returns the missing letter in the array.
// You will always get an valid array. And it will be always exactly one letter be missing. The length of the array will always be at least 2.
// The array will always contain letters in only one case.
// Example:
// ['a','b','c','d','f'] -> 'e'
// ['O','Q','R','S'] -> 'P'
// (Use the English alphabet with 26 letters!)
// Have fun coding it and please don't forget to vote and rank this kata! :-)
// I have also created other katas. Take a look if you enjoyed this kata!

// Answer :
package soal


func FindMissingLetter(chars []rune) (res rune) {

	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	abjad := map[rune]int{}

	for _, t := range text {
		abjad[t] = int(t)
	}


	for i := 0; i < len(chars) - 1; i++{
		intChars := int(chars[i])
		nextIntChars := int(chars[i + 1])

		if nextIntChars != intChars + 1{
			res = rune(abjad[chars[i + 1] - 1])
		}
	} 

	return 
}