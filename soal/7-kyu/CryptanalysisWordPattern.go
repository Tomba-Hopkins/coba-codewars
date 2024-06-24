// Question :

// In cryptanalysis, words patterns can be a useful tool in cracking simple ciphers.
// A word pattern is a description of the patterns of letters occurring in a word, where each letter is given an integer code in order of appearance. So the first letter is given the code 0, and second is then assigned 1 if it is different to the first letter or 0 otherwise, and so on.
// As an example, the word "hello" would become "0.1.2.2.3". For this task case-sensitivity is ignored, so "hello", "helLo" and "heLlo" will all return the same word pattern.
// Your task is to return the word pattern for a given word. All words provided will be non-empty strings of alphabetic characters only, i.e. matching the regex "[a-zA-Z]+".

// Answer :
package soal

import (
	"strconv"
	"strings"
)

func CryptanalysisWordPattern(word string) string {

	box := map[string]int{}
	dahliat := map[string]bool{}
	low := strings.ToLower(word)


	d := 0

	for _, l := range low{
		str := string(l)
		if !dahliat[str] {
			box[str] = d
			dahliat[str] = true
			d++
		}
	}

	rilBox := []string{}

	for _, v := range low{

		newV := string(v)
		if box[newV] >= 0 {
			toStr := strconv.Itoa(box[newV])
			rilBox = append(rilBox, toStr)
		}

	}
	
	return strings.Join(rilBox, ".")
}