package views

import (
	"fmt"
	"log"
	"time"

	"github.com/Abdirahman04/novatrust/pkg/model"
	"github.com/Abdirahman04/novatrust/pkg/service"
	"github.com/Abdirahman04/novatrust/pkg/styles"
	"github.com/Abdirahman04/novatrust/pkg/utils"
)


func Signup() {
  styles.Head("sign up page")
  id := utils.Form("Customer ID") 
  fname := utils.Form("First name")
  lname := utils.Form("Last name")
  email := utils.Form("Email address")
  pass := utils.Form("Password")

  customer := model.Customer{CustomerId: id, FirstName: fname, LastName: lname, Email: email, Password: pass, DOC: time.Now()}
  res, err := service.SaveCustomer(customer)

  if err != nil {
    log.Fatal(err)
    fmt.Println("Unable to add customer!")
  }

  fmt.Println(res)
  
  Dashboard(customer)
}
