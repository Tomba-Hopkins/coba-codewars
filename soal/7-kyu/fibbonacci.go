// Question :
// Create function fib that returns n'th element of Fibonacci sequence (classic programming task).

// Answer :
package soal

func Fib(n int) int {

	a := 0
	b := 1
	c := 0

	if(n == 0) {
		return a
	}

	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}

	return b
}