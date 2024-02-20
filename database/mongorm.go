package database

import (
	"context"

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
}

// type articleAdapter struct {
// 	article ArticleAdapter
// }

// func (W *articleAdapter) InsertArticle(article Article) {
// 	W.article.insertArticleIntoDatabase(article)
// }

type ArticleMongoDb struct {
	Client *mongo.Client
}

func (AM *ArticleMongoDb) InsertArticleIntoDatabase(article Article) *mongo.InsertOneResult {
	articleCollection := AM.Client.Database("testing").Collection("article")

	insertedData, err := articleCollection.InsertOne(context.TODO(), article)

	if err != nil {
		panic(err)
	}

	return insertedData
}
