package connection

import (
	"context"
	"log"
	"nms/src/backend/env"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct{}

func (this MongoDB) newContext(sec int32) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)
}
func (this MongoDB) connect() (*mongo.Client, error) {
	ctx, cancel := this.newContext(1)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(env.MONGODB_SRV),
	)
	if err != nil {
		return nil, err
	}

	return client, err
}
func (this MongoDB) GetDatabaseNames() []string {
	client, err := this.connect()
	if err != nil {
		log.Println(err)
		return nil
	}

	ctx, cancel := this.newContext(1)
	defer func() {
		cancel()
		if err := client.Disconnect(ctx); err != nil {
			log.Println(err)
		}
	}()

	listDatabaseNames, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil
	}

	return listDatabaseNames
}
func (this MongoDB) GetDatabaseNameWithCollectionName() []map[string]interface{} {
	client, err := this.connect()
	if err != nil {
		log.Println(err)
		return nil
	}

	ctx, cancel := this.newContext(2)
	defer func() {
		cancel()
		if err := client.Disconnect(ctx); err != nil {
			log.Println(err)
		}
	}()

	listDatabaseNamesWithListColletionNames := []map[string]interface{}{}
	for _, database := range this.GetDatabaseNames() {
		listCollectionNames, err := client.Database(database).ListCollectionNames(ctx, bson.M{})
		if err != nil {
			return nil
		}

		listDatabaseNamesWithListColletionNames = append(listDatabaseNamesWithListColletionNames,
			map[string]interface{}{
				"name":        database,
				"collections": listCollectionNames,
			},
		)
	}

	return listDatabaseNamesWithListColletionNames
}
