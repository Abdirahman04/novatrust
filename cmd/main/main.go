package main

import (
	"github.com/Abdirahman04/novatrust/pkg/utils"
	"github.com/Abdirahman04/novatrust/pkg/views"
)

func main() {
  views.Homepage()
  ans := utils.Menu()
  if ans == "1\n" {
    views.Signup()
  }
}
