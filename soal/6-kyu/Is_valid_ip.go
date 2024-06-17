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

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {

	arr := strings.Split(ip, ".")

	if len(ip) == 0 {
		return false
	} else if len(arr) != 4 {
		return false
	}

	hurufJanganMasuk := "1234567890"
	khususNum := map[string]bool{}
	for _, h := range hurufJanganMasuk{
		khususNum[string(h)] = true
	}

	for _, i := range arr{
		if string(i[0]) == "0" && len(i) > 1 {
			return false
		}
		for _, h := range i {
			if !khususNum[string(h)]{
				return false
			} else if h == ' '{
				return false
			}
		}
	}


	for _, r := range arr {
		num, _ := strconv.Atoi(r)
		if num > 255 || num < 0 {
			return false
		}
	}

	return true

}