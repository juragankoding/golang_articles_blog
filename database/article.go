package database

import (
	"context"

	"example.com/chin/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Article struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Permalink   string             `bson:"permalink,omitempty"`
	Description string             `bson:"description,omitempty"`
}

type ArticleAdapter interface {
	InsertArticleIntoDatabase(article Article) *mongo.InsertOneResult
	ReadSingleArticle(id string) Article
}

type ArticleMongoDb struct {
	Client *mongo.Client
}

func (AM *ArticleMongoDb) InsertArticleIntoDatabase(article Article) *mongo.InsertOneResult {
	articleCollection := AM.Client.Database(env.GetMongoDatabase()).Collection("article")

	insertedData, err := articleCollection.InsertOne(context.TODO(), article)

	if err != nil {
		panic(err)
	}

	return insertedData
}

func (AM *ArticleMongoDb) ReadSingleArticle(id string) Article {
	articleCollection := AM.Client.Database(env.GetMongoDatabase()).Collection("article")

	article := articleCollection.FindOne(context.TODO(), bson.M{})

	articleData := Article{}

	article.Decode(&articleData)

	return articleData
}
