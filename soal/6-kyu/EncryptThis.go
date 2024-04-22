// Question :

// Acknowledgments:
// I thank yvonne-liu for the idea and for the example tests :)
// Description:
// Encrypt this!
// You want to create secret messages which can be deciphered by the Decipher this! kata. Here are the conditions:
// Your message is a string containing space separated words.
// You need to encrypt each word in the message using the following rules:
// The first letter must be converted to its ASCII code.
// The second letter must be switched with the last letter
// Keepin' it simple: There are no special characters in the input.
// Examples:
// EncryptThis("Hello") == "72olle"
// EncryptThis("good") == "103doo"
// EncryptThis("hello world") == "104olle 119drlo"

// Answer :
package soal

import (
	"strconv"
	"strings"
)

func EncryptThis(text string) string {

	if len(text) == 0 {
		return ""
	  }  
	  
	if len(text) == 1 {
		return strconv.Itoa(int(text[0]))
	  }

	res := ""
	
	sentences := strings.Split(text, " ")

	for i := 0; i < len(sentences); i++{
		kalimat := sentences[i]
		asci := strconv.Itoa(int(kalimat[0]))

		karakter := []byte(kalimat)

		if len(kalimat) > 1 {
			karakter[1], karakter[len(karakter) - 1] = karakter[len(karakter) - 1], karakter[1]
		}

		kalimat = string(karakter)

		new := asci + kalimat[1:]

		if i < len(sentences) - 1{
			res += new + " "
		} else {
			res += new
		}

	}

	return res

}