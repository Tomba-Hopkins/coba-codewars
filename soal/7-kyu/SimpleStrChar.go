// Question :

// In this Kata, you will be given a string and your task will be to return a list of ints detailing the count of uppercase letters, lowercase, numbers and special characters, as follows.
// Solve("*'&ABCDabcde12345") = [4,5,5,3].
// --the order is: uppercase letters, lowercase, numbers and special characters.
// More examples in the test cases.
// Good luck!

// Answer :
package soal

func SimpleStrChar(s string) (res []int) {

	up, low, num, other := 0, 0, 0, 0


	for i := 0; i < len(s); i++ {

		str := string(s[i])
		switch {
		case str >= "A" && str <= "Z":
			up++
		case str >= "a" && str <= "z":
			low++
		case str >= "0" && str <= "9":
			num++
		default:
			other++
		}

	}

	res = append(res, up, low, num, other)

	return

}