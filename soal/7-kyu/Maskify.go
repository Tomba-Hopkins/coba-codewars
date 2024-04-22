// Question :
// Usually when you buy something, you're asked whether your credit card number, phone number or answer to your most secret question is still correct. However, since someone could look over your shoulder, you don't want that shown on your screen. Instead, we mask it.
// Your task is to write a function maskify, which changes all but the last four characters into '#'.
// Examples (input --> output):
// "4556364607935616" --> "############5616"
//      "64607935616" -->      "#######5616"
//                "1" -->                "1"
//                 "" -->                 ""
// // "What was the name of your first pet?"
// "Skippy" --> "##ippy"
// "Nananananananananananananananana Batman!" --> "####################################man!"

// Answer :
package soal
func Maskify(word string) string {

	result := ""

	if len(word) == 1{
		return word
	} else if len(word) == 0{
		return ""
	}

	for i := 0; i < len(word); i++ {
		if i < len(word)-4 {
			result += "#"
		} else {
			result += string(word[i])
		}
	}
	return result

}

// versi js :
// function maskify(cc) {
// 	if (cc.length == 1){
// 	  return cc
// 	} else if (cc.length == 0){
// 	  return ""
// 	}
// 	let result = ""
// 	for (let i = 0; i < cc.length; i++){
// 	  if (i < cc.length - 4){
// 		result += "#"
// 	  } else {
// 		result += cc[i]
// 	  }
	  
// 	}
// 	return result
//   }