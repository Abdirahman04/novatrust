package model

import "time"

type Account struct {
  AccountId string
  CustomerId string
  AccountType string
  Balance float32
  DOC time.Time
}
