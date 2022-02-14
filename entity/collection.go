////////////////////////////////////////////////////////////////////////////////
//	collection.go  -  Feb/12/2022  -  aldebap
//
//	Database collection
////////////////////////////////////////////////////////////////////////////////

package entity

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var authorCollection *mongo.Collection
var publisherCollection *mongo.Collection
var bookCollection *mongo.Collection

//	init every database collection
func InitCollections(dbClient *mongo.Client) {
	authorCollection = dbClient.Database("Bookshelf").Collection("Author")
	publisherCollection = dbClient.Database("Bookshelf").Collection("Publisher")
	bookCollection = dbClient.Database("Bookshelf").Collection("Book")
}
