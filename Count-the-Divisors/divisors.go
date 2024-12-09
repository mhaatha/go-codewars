// https://www.codewars.com/kata/542c0f198e077084c0000c2e/train/go

package main

func Divisors(n int) int {
	result := 1

	dividedN := n / 2
	for i := 1; i <= dividedN; i++ {
		if n%i == 0 {
			result++
		}
	}

	return result
}

/*
AI Solution
func countDivisors(n int) int {
    count := 0
    for i := 1; i*i <= n; i++ {
        if n % i == 0 { // Jika n % i == 0 maka i adalah pembagi n
            if i == n / i { // Jika i == n / i (Jika i adalah akar kuadrat n)
                count += 1 // Pembagi unik (akar kuadrat)
            } else {
                count += 2 // Pasangan pembagi
            }
        }
    }
    return count
}

*/
