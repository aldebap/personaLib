////////////////////////////////////////////////////////////////////////////////
//	book.go  -  Feb/13/2022  -  aldebap
//
//	Book entity
////////////////////////////////////////////////////////////////////////////////

package store

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//	book attributes
type Book struct {
	Id    string `bson:"_id,omitempty"`
	Title string `bson:"title,omitempty"`
	//Author    []Author  `bson:"author,omitempty"`
	Publisher Publisher `bson:"publisher,omitempty"`
}

//	get all book from database
func GetAllBook() ([]Book, error) {

	//	publisher aggreggation pipeline
	lookupPublisher := bson.D{{"$lookup", bson.D{{"from", "Publisher"}, {"localField", "publisher"}, {"foreignField", "_id"}, {"as", "publisher"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$publisher"}, {"preserveNullAndEmptyArrays", false}}}}

	//	get a cursor for the query
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := bookCollection.Aggregate(ctx, mongo.Pipeline{lookupPublisher, unwindStage})
	if err != nil {
		log.Default().Printf("[1] Mongo DB error: %v", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	//	fetch data from cursor
	var bookList []Book

	err = cursor.All(ctx, &bookList)
	if err != nil {
		log.Default().Printf("[2] Mongo DB error: %v", err.Error())
		return nil, err
	}

	return bookList, nil
}
