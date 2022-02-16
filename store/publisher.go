////////////////////////////////////////////////////////////////////////////////
//	publisher.go  -  Feb/12/2022  -  aldebap
//
//	Published entity
////////////////////////////////////////////////////////////////////////////////

package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//	publisher attributes
type Publisher struct {
	Id   string `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
}

//	get all publisher from database
func GetAllPublisher() ([]Publisher, error) {

	//	get a cursor for the query
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := publisherCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	//	fetch data from cursor
	var publisherList []Publisher

	err = cursor.All(ctx, &publisherList)
	if err != nil {
		return nil, err
	}

	return publisherList, nil
}
