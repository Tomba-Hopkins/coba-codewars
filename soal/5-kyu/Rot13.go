// Question :

// How can you tell an extrovert from an introvert at NSA?
// Va gur ryringbef, gur rkgebireg ybbxf ng gur BGURE thl'f fubrf.
// I found this joke on USENET, but the punchline is scrambled. Maybe you can decipher it?
// According to Wikipedia, ROT13 is frequently used to obfuscate jokes on USENET.
// For this task you're only supposed to substitute characters. Not spaces, punctuation, numbers, etc.
// Test examples:
// "EBG13 rknzcyr." -> "ROT13 example."
// "This is my first ROT13 excercise!" -> "Guvf vf zl svefg EBG13 rkprepvfr!"

// Answer :
package soal

import "strings"

func Rot13(msg string) (res string) {

	text := "abcdefghijklmnopqrstuvwxyz"
	abjadNum := map[string]int{}
	numAbjad := map[int]string{}
	
	for index, an := range text{
		abjadNum[string(an)] = index + 1
	}

	for index, an2 := range text{
		numAbjad[index + 1] = string(an2)
	}

	for _, m := range msg{

		strM := string(m)
		strlow := strings.ToLower(strM)

		if abjadNum[strM] > 0 {

			num := abjadNum[strM]

			if num + 13 > len(abjadNum){
				num += 13 - len(abjadNum)
			} else {
				num += 13
			}

			res += numAbjad[num]
			
		} else if abjadNum[strlow] > 0 {

			num := abjadNum[strlow]

			if num + 13 > len(abjadNum){
				num += 13 - len(abjadNum)
			} else {
				num += 13
			}

			res += strings.ToUpper(numAbjad[num])
			
		} else {
			res += strM
		}
		
	}

	return
}