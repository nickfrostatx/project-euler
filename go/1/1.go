package main

import "fmt"

func triangle(n int) int {
  return n * (n + 1) / 2
}

func main() {
  var sum = 3 * triangle(999 / 3) + 5 * triangle(999 / 5) - 15 * triangle(999 / 15)
  fmt.Println(sum)
}
