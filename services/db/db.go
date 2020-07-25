package db

import (
	"context"
	"fmt"
	"services/utils"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Conn connection
type Conn struct {
	Collection *mongo.Collection
	Ctx        *context.Context
}

var (
	once sync.Once
	conn *Conn
)

func init() {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI("mongodb://root:root@127.0.0.1:27017")
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			fmt.Println(err)
			fmt.Println(utils.Trace())
		}
		ctx, _ := context.WithTimeout(context.Background(), 100000*time.Second)
		col := client.Database("Emails").Collection("email")
		conn = &Conn{
			Collection: col,
			Ctx:        &ctx,
		}
	})
}

// GetDb connection singleton
func GetDb() *Conn {
	return conn
}
