// Question :

// altERnaTIng cAsE <=> ALTerNAtiNG CaSe
// Define String.prototype.toAlternatingCase (or a similar function/method such as to_alternating_case/toAlternatingCase/ToAlternatingCase in your selected language; see the initial solution for details) such that each lowercase letter becomes uppercase and each uppercase letter becomes lowercase. For example:
// ToAlternatingCase("hello world"); // => "HELLO WORLD"
// ToAlternatingCase("HELLO WORLD"); // => "hello world"
// ToAlternatingCase("hello WORLD"); // => "HELLO world"
// ToAlternatingCase("HeLLo WoRLD"); // => "hEllO wOrld"
// ToAlternativeCase("12345"); // => "12345" (Non-alphabetical characters are unaffected)
// ToAlternativeCase("1a2b3c4d5e"); // => "1A2B3C4D5E"
// ToAlternativeCase("String.prototype.toAlternatingCase"); // => "sTRING.PROTOTYPE.TOaLTERNATINGcASE"
// As usual, your function/method should be pure, i.e. it should not mutate the original string.

// Answer :
package soal

import "strings"

func ToAlternatingCase(str string) (res string) {

	normal := map[string]bool{}
	huruf := "abcdefghijklmnopqrstuvwxyz"
	for _, r := range huruf {
		normal[string(r)] = true
	}

	for _, s := range str {
		if normal[string(s)] {
			cap := strings.ToUpper(string(s))
			res += cap
		} else {
			nor := strings.ToLower(string(s))
			res += nor
		}
	}

	return 
}