package mongodb

import (
	"context"
	"eta-of-taipeimetro/configuration"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx                     context.Context
	cancel                  context.CancelFunc
	client                  *mongo.Client
	DB                      *mongo.Database
	LineTransferCollection  *mongo.Collection
	S2STravelTimeCollection *mongo.Collection
)

func Initialize() {
	ctx, cancel = context.WithCancel(context.Background())

	// Variables
	address := configuration.Conf.MongoDB_Address
	port := configuration.Conf.MongoDB_Port

	// For Role-Based Access Control is set in MongoDB Server.
	user := ""
	pwd := ""

	dbName := "testDB"

	// Setup connection
	initMongoDBClients(user, pwd, address, port, dbName)
}

func initMongoDBClients(user string, pwd string, address string, port int, dbName string) {
	// MongoDB uri
	var uri string
	if user != "" && pwd != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%v", user, pwd, address, port)
	}
	uri = fmt.Sprintf("mongodb://%s:%v", address, port)

	client = setupConnection(context.TODO(), uri, user, pwd)
	DB := client.Database(dbName)
	LineTransferCollection = DB.Collection("LineTransfer")
	S2STravelTimeCollection = DB.Collection("S2STravelTime")
}

func setupConnection(ctx context.Context, uri string, user string, pwd string) *mongo.Client {
	// Create a new ClientOptions instance.
	clientOpts := options.Client().ApplyURI(uri)

	if user != "" && pwd != "" {
		// Create Credential instance.
		auth := options.Credential{
			Username: user,
			Password: pwd,
		}
		clientOpts.SetAuth(auth)
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)

	// Create a new client and connect to the server.
	log.Printf("[Info] Connect to MongoDB : %s.\n", uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(err)
	}

	defer cancel()

	// Verify connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("[Info] Successfully connected.")

	return client
}

func CloseConnection() {
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
	log.Println("[Info] Connection to MongoDB closed.")
}
