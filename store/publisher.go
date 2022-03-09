////////////////////////////////////////////////////////////////////////////////
//	publisher.go  -  Feb/12/2022  -  aldebap
//
//	Published entity
////////////////////////////////////////////////////////////////////////////////

package store

import (
	"context"
	"personaLib/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	publisher attributes
type Publisher struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}

//	create a new model Publisher from store document publisher
func PublisherFromDocument(publisher Publisher) *model.Publisher {
	return &model.Publisher{
		Id:   publisher.Id.Hex(),
		Name: publisher.Name,
	}
}

//	add publisher to collection
func AddPublisher(publisher *model.Publisher) (*model.Publisher, error) {

	var newPublisher Publisher

	newPublisher.Id = primitive.NewObjectID()
	newPublisher.Name = publisher.Name

	//	add the publisher to the collection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := publisherCollection.InsertOne(ctx, newPublisher)
	if err != nil {
		return nil, err
	}

	return PublisherFromDocument(newPublisher), nil
}

//	get publisher by ID from collection
func GetPublisherByID(Id *model.ID) (*model.Publisher, error) {

	objectId, err := primitive.ObjectIDFromHex(string(*Id))
	if err != nil {
		return nil, err
	}

	//	find the publisher by Id
	var publisher Publisher

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = publisherCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&publisher)
	if err != nil {
		return nil, err
	}

	return PublisherFromDocument(publisher), nil
}

//	get all publishers from collection
func GetAllPublisher() ([]model.Publisher, error) {

	//	get a cursor for the query
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := publisherCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	//	fetch data from cursor
	var publisherList []Publisher
	var resultPublisherList []model.Publisher

	err = cursor.All(ctx, &publisherList)
	if err != nil {
		return nil, err
	}

	for _, item := range publisherList {
		resultPublisherList = append(resultPublisherList, *PublisherFromDocument(item))
	}

	return resultPublisherList, nil
}

//	update publisher by ID in the collection
func UpdatePublisher(Id *model.ID, publisher *model.Publisher) error {

	objectId, err := primitive.ObjectIDFromHex(string(*Id))
	if err != nil {
		return err
	}

	var replacePublisher Publisher

	replacePublisher.Id = objectId
	replacePublisher.Name = publisher.Name

	//	update the publisher in the collection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = publisherCollection.ReplaceOne(ctx, bson.M{"_id": objectId}, replacePublisher)
	if err != nil {
		return err
	}

	return nil
}

//	delete the publisher by ID from collection
func DeletePublisher(Id *model.ID) error {

	//	delete the publisher by Id
	id, err := primitive.ObjectIDFromHex(string(*Id))
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = publisherCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
