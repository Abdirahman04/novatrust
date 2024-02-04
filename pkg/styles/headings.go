package styles

import (
  "fmt"
  "strings"
)

func Head(text string) {
  fmtext := strings.ToUpper(text)
  fmt.Printf(">>>>>     %v     <<<<<\n", fmtext)
}
