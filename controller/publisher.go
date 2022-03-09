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

	"github.com/gorilla/mux"
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

//	add publisher API
func AddPublisher(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	check for "json" content type
	contentType := httpRequest.Header.Get("Content-type")
	if "application/json" != contentType {
		httpResponse.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	//	check for non empty content length
	contentLength := httpRequest.Header.Get("Content-Length")
	if "" == contentLength {
		httpResponse.WriteHeader(http.StatusLengthRequired)
		return
	}

	if 0 == httpRequest.ContentLength {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	fetch request payload
	var requestData publisherRequest

	err := json.NewDecoder(httpRequest.Body).Decode(&requestData)
	if nil != err {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	publisher := model.NewPublisher(requestData.Name)
	if !publisher.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	add the publisher to database
	insertedPublisher, err := store.AddPublisher(publisher)
	if nil != err {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	fill response payload
	var responseData = publisherResponse{}

	responseData.ID = insertedPublisher.Id
	responseData.Name = insertedPublisher.Name

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusCreated)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	get publisher by Id API
func GetPublisher(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)
	id := model.NewID(vars["id"])

	if !id.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	get the publisher by Id from database
	var publisher *model.Publisher

	publisher, err := store.GetPublisherByID(id)
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	var responseData = publisherResponse{}

	responseData.ID = publisher.Id
	responseData.Name = publisher.Name

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
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

//	patch publisher by Id API
func PatchPublisher(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)
	id := model.NewID(vars["id"])

	if !id.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	fetch request payload
	var requestData publisherRequest

	err := json.NewDecoder(httpRequest.Body).Decode(&requestData)
	if nil != err {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(requestData.ID) > 0 && requestData.ID != string(*id) {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	publisher := model.NewPublisher(requestData.Name)
	if !publisher.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	update the publisher by Id in the database
	err = store.UpdatePublisher(id, publisher)
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	var responseData = publisherResponse{}

	responseData.ID = string(*id)
	responseData.Name = publisher.Name

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	delete publisher by Id API
func DeletePublisher(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)
	id := model.NewID(vars["id"])

	if !id.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	delete the publisher by Id from database
	err := store.DeletePublisher(id)
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
}
