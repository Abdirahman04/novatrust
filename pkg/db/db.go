package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
  clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

  client, err := mongo.Connect(context.TODO(), clientOptions)

  if err != nil {
    log.Fatal(err)
  }

  err = client.Ping(context.TODO(), nil)

  if err != nil {
    log.Fatal(err)
  }

  return client, nil
}

func Disconnect(client *mongo.Client) error {
  if client != nil {
    err := client.Disconnect(context.Background())
    return err
  }
  return nil
}
