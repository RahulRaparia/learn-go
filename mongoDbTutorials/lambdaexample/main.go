package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventInput struct {
	Limit int64 `json:"limit"`
}

type Movie struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id"`
	Title string 		   `bson:"title" json:"title"`
	Year  int32             `bson:"year" json:"year"`
}

var client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("mongodb+srv://rahulraparia:qwerty123@cluster0.nyn9bvv.mongodb.net/?retryWrites=true&w=majority")))

func HandleRequest(ctx context.Context, input EventInput) ([]Movie, error) {
	//the connection to the database could have failed, so we're catching it here.
	if err != nil {
		return nil, err
	}

	collection := client.Database("sample_mflix").Collection("movies")

	opts := options.Find()

	if input.Limit != 0 {
		opts = opts.SetLimit(input.Limit)
	}

	cursor, err := collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	var movies []Movie
	if err = cursor.All(context.Background(), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func main() {
	lambda.Start(HandleRequest)
}