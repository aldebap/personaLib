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
)

//	author attributes
type Author struct {
	Id   string `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
}

//	get all author from database
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
