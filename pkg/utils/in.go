package utils

import (
  "log"
  "bufio"
  "os"
  "fmt"
)

func Menu() string {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(">>>>> ")
  name, err := reader.ReadString('\n')

  if err != nil {
    log.Fatal(err)
  }

  return name
}
