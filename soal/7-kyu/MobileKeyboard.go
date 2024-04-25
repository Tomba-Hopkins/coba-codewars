// Question :

// Do you remember the old mobile display keyboards? Do you also remember how inconvenient it was to write on it?
// Well, here you have to calculate how many keystrokes you have to do for a specific word.
// This is the layout:
// Given a string, return the number of keystrokes necessary to type it. You can assume that the input will entirely consist of characters included in the mobile layout (lowercase letters, digits, and the symbols * and #).
// Examples
// "*#"       =>  2 (1 + 1)
// "123"      =>  3 (1 + 1 + 1)
// "abc"      =>  9 (2 + 3 + 4)
// "codewars" => 26 (4 + 4 + 2 + 3 + 2 + 2 + 4 + 5)

// Answer :
package soal

func MobileKeyboard(str string) (res int) {

	nokia := map[rune]int{}
	satu := "0123456789*#"
	dua := "adgjmptw"
	tiga := "behknqux"
	empat := "cfilorvy"
	lima := "sz"

	for _, s := range satu {
		nokia[s] = 1
	}
	for _, d := range dua {
		nokia[d] = 2
	}
	for _, t := range tiga {
		nokia[t] = 3
	}
	for _, e := range empat {
		nokia[e] = 4
	}
	for _, l := range lima {
		nokia[l] = 5
	}

	for _, result := range str {
		res += nokia[result]
	}

	return
}
