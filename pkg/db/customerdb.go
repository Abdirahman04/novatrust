package db

import (
	"context"
	"fmt"
	"log"

	"github.com/Abdirahman04/novatrust/pkg/model"
)

func AddCustomer(customer model.Customer) (string, error) {
  client, err := Connect()

  if err != nil {
    log.Fatal(err)
    return "", err
  }

  collection := client.Database("novatrust").Collection("customers")

  insertResult, err := collection.InsertOne(context.TODO(), customer)

  if err != nil {
    log.Fatal(err)
    return "", err
  }

  msg := fmt.Sprintln("Inserted a single document,", insertResult.InsertedID)
  return msg, nil
}
