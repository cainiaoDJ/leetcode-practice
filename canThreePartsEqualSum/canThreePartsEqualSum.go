package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
	avg := 0
	for _, v := range A {
		avg += v
	}
	if avg%3 != 0 {
		return false
	}
	avg /= 3
	l := len(A)
	head, tail := 0, l-1
	sum1, sum3 := 0, 0
	for head < tail {
		if (head != 0 && sum1 != avg) || head == 0 {
			sum1 += A[head]
			head++
		}
		if (tail != l-1 && sum3 != avg) || tail == l-1 {
			sum3 += A[tail]
			tail--
		}
		if sum1 == sum3 && sum1 == avg && head <= tail {
			return true
		}

	}

	return false
}

func main() {
	a := []int{18, 12, -18, 18, -19, -1, 10, 10}
	fmt.Println(canThreePartsEqualSum(a))
}
