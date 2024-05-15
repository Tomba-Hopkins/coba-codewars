// Question :

// Task
// Given a string str, reverse it and omit all non-alphabetic characters.
// Example
// For str = "krishan", the output should be "nahsirk".
// For str = "ultr53o?n", the output should be "nortlu".
// Input/Output
// [input] string str
// A string consists of lowercase latin letters, digits and symbols.
// [output] a string

// Answer :
package soal

func ReverseLetters(s string) (res string) {

	txt := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	abjad := map[string]bool{}

	for _, r := range txt{
		abjad[string(r)] = true
	}

	for i := len(s) - 1; i >= 0; i-- {
		if abjad[string(s[i])]{
			res += string(s[i])
		}
	}
	return
}