package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


type Store interface {
	Insert(context.Context,*Event)error
	UpdateStatus(context.Context,*Event)error
}


type MongoStore struct {
	collection *mongo.Database
}

func GetStore(ctx context.Context)(*MongoStore,error){
	godotenv.Load()
	url:=os.Getenv("MONGO_URL_EVENT_DOCUMENT")
	client,err:=mongo.Connect(ctx,options.Client().ApplyURI(url))
	if err!=nil{
		return nil,err
	}
	if err:=client.Ping(ctx,readpref.Primary());err!=nil{
		return nil,err
	}
	database:=client.Database("Event-Document")
	return &MongoStore{
		collection: database,
	},nil
}


func(store *MongoStore)Insert(ctx context.Context,event *Event)error{
	coll:=store.collection.Collection("Event-Document")
	_,err:=coll.InsertOne(ctx,event)
	if err!=nil{
		return err
	}
	return nil
}

func(store *MongoStore)UpdateStatus(ctx context.Context,event *Event)error{
	coll:=store.collection.Collection("Event-Document")
	filter:=bson.M{
		"_id":event.Id,
	}
	update:=bson.M{
		"$set":bson.M{
			"status":"processed",
		},
	}
	_,err:=coll.UpdateOne(ctx,filter,update)
	if err!=nil{
		return err
	}
	return nil

}

