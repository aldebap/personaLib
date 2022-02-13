////////////////////////////////////////////////////////////////////////////////
//	controller.go  -  Feb/11/2022  -  aldebap
//
//	personLib routes controller
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"net/http"

	"personaLib/entity"
)

//	get author response
type getAuthorResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//	fetch all authors response
type getAllAuthorResponse struct {
	Author []getAuthorResponse `json:"author"`
}

//	get all authors API
func getAllAuthors(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	get all authors from database
	var authorList []entity.Author

	authorList, err := entity.GetAllAuthor()
	if err != nil {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	fill response payload
	var responseData = getAllAuthorResponse{}

	for _, item := range authorList {

		var authorData = getAuthorResponse{}

		authorData.ID = item.Id
		authorData.Name = item.Name

		responseData.Author = append(responseData.Author, authorData)
	}

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	get publisher response
type getPublisherResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//	fetch all publishers response
type getAllPublisherResponse struct {
	Author []getPublisherResponse `json:"publisher"`
}

//	get all publishers API
func getAllPublishers(httpResponse http.ResponseWriter, httpRequest *http.Request) {

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

		responseData.Author = append(responseData.Author, publisherData)
	}

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}
