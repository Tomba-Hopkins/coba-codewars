// Question :

// If　a = 1, b = 2, c = 3 ... z = 26
// Then l + o + v + e = 54
// and f + r + i + e + n + d + s + h + i + p = 108
// So friendship is twice as strong as love :-)
// Your task is to write a function which calculates the value of a word based off the sum of the alphabet positions of its characters.
// The input will always be made of only lowercase letters and will never be empty.

// Answer :
package soal

func WordsToMarks(s string) int {
	result := 0
	abjad := map[string]int{}
	text := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(text); i++ {
		// abjad[string(text[i])]++ tadi gini
		abjad[string(text[i])] = i + 1
	}

	for i := 0; i < len(s); i++{
		notByteAnymore := string(s[i])
		result += abjad[notByteAnymore]
	}

	return result
}