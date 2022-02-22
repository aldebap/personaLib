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

//	create a new model Author from store document author
func FromDocument(author Author) *model.Author {
	return &model.Author{
		Id:   author.Id.Hex(),
		Name: author.Name,
	}
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

	return FromDocument(newAuthor), nil
}

//	get author by ID from collection
func GetAuthorByID(Id string) (*model.Author, error) {

	objectId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, err
	}

	//	find the author by Id
	var author Author

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = authorCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&author)
	if err != nil {
		return nil, err
	}

	return FromDocument(author), nil
}

//	get all authors from collection
func GetAllAuthor() ([]model.Author, error) {

	//	get a cursor for the query
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := authorCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	//	fetch data from cursor
	var authorList []Author
	var resultAuthorList []model.Author

	err = cursor.All(ctx, &authorList)
	if err != nil {
		return nil, err
	}

	for _, item := range authorList {
		resultAuthorList = append(resultAuthorList, *FromDocument(item))
	}

	return resultAuthorList, nil
}

//	update author by ID in the collection
func UpdateAuthor(Id string, author *model.Author) error {

	objectId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return err
	}

	var replaceAuthor Author

	replaceAuthor.Id = objectId
	replaceAuthor.Name = author.Name

	//	update the author in the collection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = authorCollection.ReplaceOne(ctx, bson.M{"_id": objectId}, replaceAuthor)
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
