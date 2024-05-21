// Question :

// Introduction
// A grille cipher was a technique for encrypting a plaintext by writing it onto a sheet of paper through a pierced sheet (of paper or cardboard or similar). The earliest known description is due to the polymath Girolamo Cardano in 1550. His proposal was for a rectangular stencil allowing single letters, syllables, or words to be written, then later read, through its various apertures. The written fragments of the plaintext could be further disguised by filling the gaps between the fragments with anodyne words or letters. This variant is also an example of steganography, as are many of the grille ciphers.
// https://en.wikipedia.org/wiki/Grille_(cryptography)

// Task
// Write a function that accepts two inputs: message and code and returns hidden message decrypted from message using the code.
// The code is a nonnegative integer and it decrypts in binary the message.

// message : abcdef
// code    : 000101
// ----------------
// result  : df

// Answer :
package soal

import "strconv"

func Grille(message string, code int) (res string) {

	bin := strconv.FormatInt(int64(code), 2)
	newBin := ""

	if len(message) < len(bin){
		temp := ""
		for i := len(bin) - 1; i >= len(bin) - len(message); i--{
			temp += string(bin[i])
		}
		for i := len(temp) - 1; i >= 0; i--{
			newBin += string(temp[i])
		}
	} else if len(bin) < len(message){
		for i := 0; i < len(message) - len(bin); i++{
			newBin += "0"
		}
		newBin += bin
	}


	for i := 0; i < len(newBin); i++{
		if newBin[i] == '1'{
			res += string(message[i])
		}
	}

	
	return 
}