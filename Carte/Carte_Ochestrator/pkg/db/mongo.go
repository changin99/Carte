package db

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "Carte_Orchestrator/config"
)

var Client *mongo.Client

// InitMongoDB는 MongoDB에 연결합니다.
func InitMongoDB() {
    clientOptions := options.Client().ApplyURI(config.Config.Database.URI)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    Client = client
    log.Println("Connected to MongoDB!")
}
