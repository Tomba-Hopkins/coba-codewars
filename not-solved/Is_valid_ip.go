// Question :
// Write an algorithm that will identify valid IPv4 addresses in dot-decimal format. IPs should be considered valid if they consist of four octets, with values between 0 and 255, inclusive.
// Valid inputs examples:
// Examples of valid inputs:
// 1.2.3.4
// 123.45.67.89
// Invalid input examples:
// 1.2.3
// 1.2.3.4.5
// 123.456.78.90
// 123.045.067.089
// Notes:
// Leading zeros (e.g. 01.02.03.04) are considered invalid
// Inputs are guaranteed to be a single string

// Answer :
package soal

import "strconv"

// jawaban ku sebelumnya :
func Is_valid_ip(ip string) bool {
	dot := "."
	dot1 := 0
	dot2 := 0
	dot3 := 0
	for i := 0; i < len(ip); i++{
		if string(ip[i]) == dot {
			dot1 = i
			break
		}
	}
	for j := dot1 + 1; j < len(ip); j++{
		if string(ip[j]) == dot {
			dot2 = j
			break
		}
	}
	for k := dot2 + 1; k < len(ip); k++{
		if string(ip[k]) == dot {
			dot3 = k
			break
		}
	}
	for _, cek := range ip {
		if string(cek) == " "{
			return false
		}
		if ip == "0.0.0.0"{
			return true
		}
		if string(ip[0]) == "0"|| string(ip[dot1 + 1]) == "0" || string(ip[dot2 + 1]) == "0" || string(ip[dot3 + 1]) == "0" {
			return false
		}
	}
	octat1 := ip[:dot1]
	res1,_ := strconv.Atoi(octat1)
	octat2 := ip[dot1 + 1:dot2]
	res2,_ := strconv.Atoi(octat2)
	octat3 := ip[dot2 + 1:dot3]
	res3,_ := strconv.Atoi(octat3)
	octat4 := ip[dot3 + 1:]
	res4,_ := strconv.Atoi(octat4)
	if res1 <= 255 && res2 <= 255 && res3 <= 255 && res4 <= 255 {
		return true
	} else {
		return false
	}
}

// WRONG