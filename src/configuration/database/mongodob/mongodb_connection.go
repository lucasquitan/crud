package mongodob

import (
	"context"

	"github.com/lucasquitan/crud/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

func InitConnection() {
	ctx := context.Background()
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logger.Error("Failed to connect to MongoDB Database", err, zap.String("journey", "mongodbConnection"))
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	defer logger.Info("MongoDB connected successfully", zap.String("journey", "mongodbConnection"))
}