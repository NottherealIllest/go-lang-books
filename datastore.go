package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Datastore struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewDatastore(ctx context.Context, mongoURI string) *Datastore {
	option := options.Client()

	client, err := mongo.NewClient(option.ApplyURI(mongoURI))
	if err != nil {
		log.Panic("Unable to create mongo client: %v", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Panic("Unable to connect mongo client: %v", err)
	}

	database := client.Database("book_store")
	collection := database.Collection("books")
	datastore := &Datastore{
		client:     client,
		database:   database,
		collection: collection,
	}

	return datastore
}

func (datastore *Datastore) CreateBook(ctx context.Context, book *Book) error {
	book.ID = primitive.NewObjectID()
	book.CreatedAt = time.Now()

	_, err := datastore.collection.InsertOne(ctx, book)
	return err
}

func (datastore *Datastore) UpdateBook(ctx context.Context, id string, book *Book) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = datastore.collection.UpdateOne(ctx, bson.M{"_id": ID}, bson.M{"$set": book})
	return err
}

func (datastore *Datastore) GetBook(ctx context.Context, id string) (*Book, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	book := &Book{}
	err = datastore.collection.FindOne(ctx, primitive.M{"_id": ID}).Decode(book)
	return book, err
}

func (datastore *Datastore) FindBooks(ctx context.Context) ([]*Book, error) {

	findOptions := options.Find()
	findOptions.SetLimit(100)

	books := []*Book{}
	cursor, err := datastore.collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for cursor.Next(ctx) {
		var elem *Book
		err = cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		books = append(books, elem)
	}

	err = cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (datastore *Datastore) DeleteBook(ctx context.Context, id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res, err := datastore.collection.DeleteOne(ctx, primitive.M{"_id": ID})
	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return fmt.Errorf("unexpected number of books deleted: %v", res.DeletedCount)
	}

	return nil
}