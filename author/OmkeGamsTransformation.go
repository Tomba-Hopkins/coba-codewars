package author

// Name: OmkeGams Transformation

// Level: 8 kyu

// Discipline: algorithm

// Tags: crypthography

// Desc:
// Implement a function omkeGas Transformation that transforms a given word based on the following rules:
// -> Every consonant from the set 'p', 'b', and 'f' is replaced with the letter 'v'.
// -> If a vowel ('a', 'i', 'u', 'e', 'o') appears at an odd index
// (starting from 0), insert the letter 'm' after the vowel.
// -> Vowels at even indices remain unchanged.
// -> Return the transformed word.
// Example :
// oke gas -> omke gams
// rispek -> rimsvek
// apple -> amvlle

// Complete Solution :
func OmkeGamsTransform(word string) (r string) {
	d := true
	vowel := map[string]bool{}
	v := "aiueo"
	vlan := map[string]bool{}
	lan := "pbf"
	for _, l := range lan {
		vlan[string(l)] = true
	}
	for _, k := range v {
		vowel[string(k)] = true
	}
	for _, w := range word {
		strW := string(w)
		if vowel[strW] && d {
			d = false
			r += strW
			r += "m"
			continue
		} else if vowel[strW] {
			d = true
			r += strW
			continue
		} else if vlan[strW] {
			r += "v"
			continue
		}
		r += strW
	}
	return
}

// Initial solution :
// func omkeGamsTransform(word string) string {
// }

// Test Case :
// import (
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	. "codewarrior/kata"
//   )
//   var _ = Describe("omkeGamsTransform", func() {
// 	It("should return 'godvlan' for input 'godplan'", func() {
// 	  Expect(omkeGamsTransform("godplan")).To(Equal("godvlan"))
// 	})
// 	It("should return 'rimsvek' for input 'rispek'", func() {
// 	  Expect(omkeGamsTransform("rispek")).To(Equal("rimsvek"))
// 	})
// 	It("should return 'omke gams' for input 'oke gas'", func() {
// 	  Expect(omkeGamsTransform("oke gas")).To(Equal("omke gams"))
// 	})
//   })