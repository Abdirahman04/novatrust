package utils

import (
  "log"
  "bufio"
  "os"
)

func Menu() string {
  reader := bufio.NewReader(os.Stdin)
  name, err := reader.ReadString('\n')

  if err != nil {
    log.Fatal(err)
  }

  return name
}
