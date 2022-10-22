package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := connect()
	if err != nil {
		panic(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	fmt.Println(err)
}

func connect() (*mongo.Client, error) {

	connectURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", "vna_tracking", "vna_tracking", "mongodb", "27017", "vna_tracking")
	opt := options.Client().ApplyURI(connectURI)
	opt.SetServerSelectionTimeout(15 * time.Second)
	opt.SetMaxConnIdleTime(10 * time.Second)
	client, err := mongo.Connect(context.Background(), opt, nil)
	if err != nil {
		fmt.Println(err, client)
		return nil, err
	}

	return client, nil
}
