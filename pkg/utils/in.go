package utils

import (
  "fmt"
)

var p = fmt.Print

func input() string {
  var name string
  fmt.Scanln(&name)
  return name
}

func Menu() string {
  p(">>>>> ")
  return input()
}

func Form(txt string) string {
  p("...", txt, ":   ")
  return input()
}
