////////////////////////////////////////////////////////////////////////////////
//	author.go  -  Feb/11/2022  -  aldebap
//
//	Author entity
////////////////////////////////////////////////////////////////////////////////

package store

import (
	"context"
	"time"

	"personaLib/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	author attributes
type Author struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}

//	add author to collection
func AddAuthor(author *model.Author) (*model.Author, error) {

	var newAuthor Author

	newAuthor.Id = primitive.NewObjectID()
	newAuthor.Name = author.Name

	//	add the author to the collection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := authorCollection.InsertOne(ctx, newAuthor)
	if err != nil {
		return nil, err
	}

	author.Id = newAuthor.Id.Hex()

	return author, nil
}

//	get author by ID from collection
func GetAuthorByID(Id string) (*Author, error) {

	//	find the author by Id
	var author Author

	objectId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = authorCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&author)
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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := authorCollection.ReplaceOne(ctx, bson.M{"_id": author.Id}, author)
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
