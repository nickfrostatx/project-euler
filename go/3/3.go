package main

import "fmt"

const num = 600851475143

func largestPrime(prod int) int {
  factor := 2
  for factor * factor <= prod {
    if prod % factor == 0 {
      prod /= factor
    } else {
      factor++
    }
  }
  return prod
}

func main() {
  fmt.Println(largestPrime(num))
}
