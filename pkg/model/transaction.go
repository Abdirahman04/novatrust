package model

import "time"

type Transaction struct {
  TransactionId string
  AccountId string
  TransactionType string
  Target string
  Amount float32
  DOC time.Time
}
