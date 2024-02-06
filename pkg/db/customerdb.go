package db

import (
	"context"
	"fmt"
	"log"

	"github.com/Abdirahman04/novatrust/pkg/model"
)

func AddCustomer(customer model.Customer) string {
  client, err := Connect()

  if err != nil {
    log.Fatal(err)
  }

  collection := client.Database("novatrust").Collection("customers")

  insertResult, err := collection.InsertOne(context.TODO(), customer)

  if err != nil {
    log.Fatal(err)
    return fmt.Sprintln(err)
  }

  return fmt.Sprintln("Inserted a single document,", insertResult.InsertedID)
}
