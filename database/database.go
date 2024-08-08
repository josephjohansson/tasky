package database

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type MongoSecrets struct {
    MongoUser             string `json:"mongo_user"`
    MongoPassword         string `json:"mongo_password"`
    MongoConnectionString string `json:"mongo_connection_string"`
}
var Client *mongo.Client = CreateMongoClient()


func CreateMongoClient() *mongo.Client {
    secretsFile := "/secrets/mongoSecrets"
    secretsData, err := ioutil.ReadFile(secretsFile)
    if err != nil {
        log.Fatalf("Error reading secrets file: %v", err)
    }

    var secrets MongoSecrets
    err = json.Unmarshal(secretsData, &secrets)
    if err != nil {
        log.Fatalf("Error unmarshaling secrets: %v", err)
    }
    fmt.Println("MONGO URI -> ", secrets.MongoConnectionString)

    MongoDbURI := secrets.MongoConnectionString
    client, err := mongo.NewClient(options.Client().ApplyURI(MongoDbURI))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MONGO -> ", MongoDbURI)
    return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("go-mongodb").Collection(collectionName)
}
