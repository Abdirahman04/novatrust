package utils

import (
  "log"
  "bufio"
  "os"
  "fmt"
)

var p = fmt.Print

func input() string {
  reader := bufio.NewReader(os.Stdin)
  name, err := reader.ReadString('\n')

  if err != nil {
    log.Fatal(err)
  }

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
