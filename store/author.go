////////////////////////////////////////////////////////////////////////////////
//	author.go  -  Feb/11/2022  -  aldebap
//
//	Author entity
////////////////////////////////////////////////////////////////////////////////

package store

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

//	get author by ID from collection
func GetAuthorByID(Id string) (*Author, error) {

	//	find the author by Id
	var author Author

	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = authorCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&author)
	if err != nil {
		return nil, err
	}

	return &author, nil
}

//	get all authors from collection
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

//	update author by ID in the collection
func UpdateAuthor(author Author) error {

	//	update the author in the collection
	id, err := primitive.ObjectIDFromHex(author.Id)
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = authorCollection.ReplaceOne(ctx, bson.M{"_id": id}, author)
	if err != nil {
		return err
	}

	return nil
}

//	delete the author by ID from collection
func DeleteAuthor(Id string) error {

	//	delete the author by Id
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = authorCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}