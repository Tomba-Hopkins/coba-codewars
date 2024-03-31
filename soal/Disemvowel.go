// Question :

// vowels from the trolls' comments, neutralizing the threat.
// Your task is to write a function that takes a string and return a new string with all vowels removed.
// For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".
// Note: for this kata y isn't considered a vowel.

// Answer :
package soal


func Disemvowel(comment string) (res string) {

	vowel := map[string]bool{
		"a": true,
		"i": true,
		"u": true,
		"e": true,
		"o": true,
		"A": true,
		"I": true,
		"U": true,
		"E": true,
		"O": true,
	}
	for _, c := range comment {
		n := string(c)
		if !vowel[n] {
			res += n
		}
	}
	return
}