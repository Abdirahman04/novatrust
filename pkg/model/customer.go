package model

import (
	"fmt"
	"time"
)

type Customer struct {
  CustomerId string
  FirstName string
  LastName string
  Email string
  Password string
  DOC time.Time
}

func FullName(customer Customer) string {
  return fmt.Sprintf("%s %s", customer.FirstName, customer.LastName) 
}
