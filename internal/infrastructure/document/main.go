package document

import (
	"context"

	"loan_apps/config/setting"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewDocumentConnection create document connection
func NewDocumentConnection() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(setting.DocumentURL))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	return client.Database(setting.DOCDatabase)
}

// func Collection(collection string) *mongo.Collection {
// 	return Document().Collection(collection)
// }
