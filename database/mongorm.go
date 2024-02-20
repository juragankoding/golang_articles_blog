package mongorm

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
	insertArticleIntoDatabase(article Article) *mongo.InsertOneResult
}

type articleAdapter struct {
	article ArticleAdapter
}

func (W *articleAdapter) InsertArticle(article Article) {
	W.article.insertArticleIntoDatabase(article)
}

type ArticleMongoDb struct {
	client *mongo.Client
}

func (AM *ArticleMongoDb) insertArticleIntoDatabase(article Article) *mongo.InsertOneResult {
	articleCollection := AM.client.Database("testing").Collection("article")

	insertedData, err := articleCollection.InsertOne(context.TODO(), article)

	if err != nil {
		panic(err)
	}

	return insertedData
}
