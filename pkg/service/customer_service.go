package service

import (
	"log"

	"github.com/Abdirahman04/novatrust/pkg/db"
	"github.com/Abdirahman04/novatrust/pkg/model"
)

func SaveCustomer(customer model.Customer) (string, error) {
  res, err := db.AddCustomer(customer)

  if err != nil {
    log.Fatal(err)
    return "", err
  }

  return res, nil
}
