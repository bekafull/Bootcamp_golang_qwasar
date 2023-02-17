package main

import "fmt"

func my_add(a int, b int) int {
  return a + b
}

func my_sub(a int, b int) int {
  return a - b
}

func main() {
  fmt.Println(my_add(1000, 12))
  fmt.Println(my_sub(1000, 12))
}
