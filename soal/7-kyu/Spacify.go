// Question :

// Modify the spacify function so that it returns the given string with spaces inserted between each character.
// spacify=("hello world") -> # returns "h e l l o   w o r l d

// Answer :
package soal

func Spacify(s string) (res string) {


	for i := 0; i < len(s); i++{
		res += string(s[i]) + " " 
	}

	if len(res) > 0 {
		res = res[:len(res) - 1]
	}
	
	return 
}