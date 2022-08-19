package main

import "fmt"

// fib is recursive but inefficient
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

var memo = map[int]int{0: 0, 1: 1} // with known base values

// fibMemo uses memoization - a technique in which you store the results of computational
// tasks when they are completed so that when you need them again,
// you look them up instead of re-computing.
func fibMemo(n int) int {
	_, found := memo[n]
	if !found {
		memo[n] = fibMemo(n-1) + fibMemo(n-2)
	}
	return memo[n]
}

// fibIter advances to the nth fib number using a loop.
func fibIter(n int) int {
	if n == 0 {
		return n
	}
	last, next := 0, 1
	for i := 1; i < n; i++ {
		temp := last
		last = next
		next = temp + next
	}
	return next
}

// fibSeq returns a sequence of fib values
func fibSeq(n int) []int {
	values := []int{0}
	if n > 0 {
		values = append(values, 1)
	}
	last, next := 0, 1
	for i := 1; i < n; i++ {
		temp := last
		last = next
		next = temp + next
		values = append(values, next)
	}
	return values
}

func main() {
	fmt.Println(fibSeq(50))
}
