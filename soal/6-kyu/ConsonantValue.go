// Question :

// Given a lowercase string that has alphabetic characters only and no spaces, return the highest value of consonant substrings. Consonants are any letters of the alphabet except "aeiou".
// We shall assign the following values: a = 1, b = 2, c = 3, .... z = 26.
// For example, for the word "zodiacs", let's cross out the vowels. We get: "z o d ia cs"
// -- The consonant substrings are: "z", "d" and "cs" and the values are z = 26, d = 4 and cs = 3 + 19 = 22. The highest is 26.
// solve("zodiacs") = 26
// For the word "strength", solve("strength") = 57
// -- The consonant substrings are: "str" and "ngth" with values "str" = 19 + 20 + 18 = 57 and "ngth" = 14 + 7 + 20 + 8 = 49. The highest is 57.
// For C: do not mutate input.
// More examples in test cases. Good luck!

// Answer :
package soal


func ConsonantValue(str string) (res int) {

	text := "abcdefghijklmnopqrstuvwxyz"
	nilai := map[string]int{}
	for i, t := range text {
		nilai[string(t)] = i + 1
	}

	gabole := "aiueo"
	danger := map[string]bool{}
	for _, g := range gabole {
		danger[string(g)] = true
	}

	arr := []string{}
	d1 := 0
	for i := 0; i < len(str); i++ {
		if danger[string(str[i])] {
			if d1 < i {
				arr = append(arr, str[d1:i])
			}
			d1 = i + 1
		}
	}
	if d1 < len(str) {
		arr = append(arr, str[d1:])
	}

	for _, r := range arr {
		total := 0
		for _, l := range r {
			total += nilai[string(l)]
		}
		if total > res {
			res = total
		}
	}

	return
}