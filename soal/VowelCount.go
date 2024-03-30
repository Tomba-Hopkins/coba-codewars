// Question :

// Return the number (count) of vowels in the given string.
// We will consider a, e, i, o, u as vowels for this Kata (but not y).
// The input string will only consist of lower case letters and/or spaces.

// Answer :
package soal

func GetCount(str string) (count int) {
	vokal := map[string]bool{
		"a": true,
		"i": true,
		"u": true,
		"e": true,
		"o": true,
	}

	for _, s := range str {
		if vokal[string(s)] {
			count++
		}
	}

	return count
}