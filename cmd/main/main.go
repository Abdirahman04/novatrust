package main

import (
	"github.com/Abdirahman04/novatrust/pkg/db"
	"github.com/Abdirahman04/novatrust/pkg/utils"
	"github.com/Abdirahman04/novatrust/pkg/views"
)

func main() {
  db.Connect()
  exit := false
  for !exit {
    views.Homepage()

    ans := utils.Menu()
    if ans == "1" {
      views.Signup()
    } else if ans == "2" {
      views.Signin()
    } else if ans == "3" {
      views.Help()
    } else if ans == "4" {
      exit = true
    }
  }
}
