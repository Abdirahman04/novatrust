package model

import "time"

type Customer struct {
  CustomerId string
  FirstName string
  LastName string
  Email string
  Password string
  DOC time.Time
}

func FullName(customer Customer) string {
  return customer.FirstName + " " + customer.LastName
}
