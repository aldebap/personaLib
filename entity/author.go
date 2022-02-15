////////////////////////////////////////////////////////////////////////////////
//	author.go  -  Feb/11/2022  -  aldebap
//
//	Author entity
////////////////////////////////////////////////////////////////////////////////

package entity

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	author attributes
type Author struct {
	Id   string `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
}

//	add author to collection
func AddAuthor(author Author) (*Author, error) {

	var newAuthor Author

	newAuthor.Name = author.Name

	//	add the author to the collection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	insertResult, err := authorCollection.InsertOne(ctx, newAuthor)
	if err != nil {
		return nil, err
	}

	newAuthor.Id = insertResult.InsertedID.(string)

	return &newAuthor, nil
}

//	get author by ID from database
func GetAuthorByID(Id string) (*Author, error) {

	//	get a cursor for the query
	var author Author

	id, _ := primitive.ObjectIDFromHex(Id)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := authorCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&author)
	if err != nil {
		return nil, err
	}

	return &author, nil
}

//	get all authors from database
func GetAllAuthor() ([]Author, error) {

	//	get a cursor for the query
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := authorCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	//	fetch data from cursor
	var authorList []Author

	err = cursor.All(ctx, &authorList)
	if err != nil {
		return nil, err
	}

	return authorList, nil
}
