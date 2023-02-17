package main

import "fmt"

func main() {
  var a = 10
  var b = 9
  var c = 11
  var d = 10
  var y = 9
  var z = 11

  if (a > b) && (a < c) && (a == d) {
    fmt.Println("a is bigger than b AND smaller than c AND equal to d")
  }
  if (z > a) || (y > a) {
    fmt.Println("z OR y are bigger than a")
  }
}
