////////////////////////////////////////////////////////////////////////////////
//	controller.go  -  Feb/11/2022  -  aldebap
//
//	personLib routes controller
////////////////////////////////////////////////////////////////////////////////

package controller

import (
	"encoding/json"
	"net/http"

	"personaLib/entity"
)

//	get publisher response
type getPublisherResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//	fetch all publishers response
type getAllPublisherResponse struct {
	Publisher []getPublisherResponse `json:"publisher"`
}

//	get all publishers API
func GetAllPublishers(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	get all publishers from database
	var publisherList []entity.Publisher

	publisherList, err := entity.GetAllPublisher()
	if err != nil {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	fill response payload
	var responseData = getAllPublisherResponse{}

	for _, item := range publisherList {

		var publisherData = getPublisherResponse{}

		publisherData.ID = item.Id
		publisherData.Name = item.Name

		responseData.Publisher = append(responseData.Publisher, publisherData)
	}

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	get book response
type getBookResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

//	fetch all books response
type getAllBookResponse struct {
	Book []getBookResponse `json:"book"`
}

//	get all books API
func GetAllBooks(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	get all books from database
	var bookList []entity.Book

	bookList, err := entity.GetAllBook()
	if err != nil {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	fill response payload
	var responseData = getAllBookResponse{}

	for _, item := range bookList {

		var bookData = getBookResponse{}

		bookData.ID = item.Id
		bookData.Title = item.Title
		bookData.Publisher = item.Publisher.Name

		responseData.Book = append(responseData.Book, bookData)
	}

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}
