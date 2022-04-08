////////////////////////////////////////////////////////////////////////////////
//	book.go  -  Feb/13/2022  -  aldebap
//
//	Book entity
////////////////////////////////////////////////////////////////////////////////

package store

import (
	"context"
	"personaLib/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	book attributes
type Book struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title,omitempty"`
	//Author    []Author  `bson:"author,omitempty"`
	//Publisher Publisher `bson:"publisher,omitempty"`
}

//	create a new model Book from store document book
func BookFromDocument(book Book) *model.Book {
	return &model.Book{
		Id:    book.Id.Hex(),
		Title: book.Title,
	}
}

//	add book to collection
func AddBook(book *model.Book) (*model.Book, error) {

	var newBook Book

	newBook.Id = primitive.NewObjectID()
	newBook.Title = book.Title

	//	add the book to the collection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := bookCollection.InsertOne(ctx, newBook)
	if err != nil {
		return nil, err
	}

	return BookFromDocument(newBook), nil
}

//	get book by ID from collection
func GetBookByID(Id *model.ID) (*model.Book, error) {

	objectId, err := primitive.ObjectIDFromHex(string(*Id))
	if err != nil {
		return nil, err
	}

	//	find the book by Id
	var book Book

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = bookCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&book)
	if err != nil {
		return nil, err
	}

	return BookFromDocument(book), nil
}

//	get all book from database
func GetAllBook() ([]model.Book, error) {

	/*
		//	this query does a join with publisher Collection

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
	*/

	//	get a cursor for the query
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := bookCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	//	fetch data from cursor
	var bookList []Book
	var resultBookList []model.Book

	err = cursor.All(ctx, &bookList)
	if err != nil {
		return nil, err
	}

	for _, item := range bookList {
		resultBookList = append(resultBookList, *BookFromDocument(item))
	}

	return resultBookList, nil
}

//	update book by ID in the collection
func UpdateBook(Id *model.ID, book *model.Book) error {

	objectId, err := primitive.ObjectIDFromHex(string(*Id))
	if err != nil {
		return err
	}

	var replaceBook Book

	replaceBook.Id = objectId
	replaceBook.Title = book.Title

	//	update the book in the collection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = bookCollection.ReplaceOne(ctx, bson.M{"_id": objectId}, replaceBook)
	if err != nil {
		return err
	}

	return nil
}

//	delete the book by ID from collection
func DeleteBook(Id *model.ID) error {

	//	delete the book by Id
	id, err := primitive.ObjectIDFromHex(string(*Id))
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = bookCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
