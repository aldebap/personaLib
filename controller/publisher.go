////////////////////////////////////////////////////////////////////////////////
//	publisher.go  -  Mar/07/2022  -  aldebap
//
//	personLib controller - publisher API
////////////////////////////////////////////////////////////////////////////////

package controller

import (
	"encoding/json"
	"net/http"
	"personaLib/model"
	"personaLib/store"
)

//	publisher request
type publisherRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

//	publisher response
type publisherResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

//	publisher's list response
type publisherListResponse struct {
	Publisher []publisherResponse `json:"publisher"`
}

//	get all publishers API
func GetAllPublishers(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	get all publishers from database
	var publisherList []model.Publisher

	publisherList, err := store.GetAllPublisher()
	if err != nil {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	TODO: implement pagination

	//	fill response payload
	var responseData = publisherListResponse{}

	for _, item := range publisherList {

		var publisher = publisherResponse{}

		publisher.ID = item.Id
		publisher.Name = item.Name

		responseData.Publisher = append(responseData.Publisher, publisher)
	}

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}
