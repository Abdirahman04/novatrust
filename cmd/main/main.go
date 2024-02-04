package main

import (
	"github.com/Abdirahman04/novatrust/pkg/utils"
	"github.com/Abdirahman04/novatrust/pkg/views"
)

func main() {
  exit := false
  for !exit {
    views.Homepage()

    ans := utils.Menu()
    if ans == "1\n" {
      views.Signup()
    } else if ans == "2\n" {
      views.Signin()
    } else if ans == "3\n" {
      views.Help()
    } else if ans == "4\n" {
      exit = true
    }
  }
}
