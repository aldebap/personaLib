////////////////////////////////////////////////////////////////////////////////
//	author.go  -  Feb/14/2022  -  aldebap
//
//	personLib controller - author API
////////////////////////////////////////////////////////////////////////////////

package controller

import (
	"encoding/json"
	"net/http"

	"personaLib/model"
	"personaLib/store"

	"github.com/gorilla/mux"
)

//	author request
type authorRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

//	author response
type authorResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

//	author's list response
type authorListResponse struct {
	Author []authorResponse `json:"author"`
}

//	add author API
func AddAuthor(httpResponse http.ResponseWriter, httpRequest *http.Request) {

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
	var requestData authorRequest

	err := json.NewDecoder(httpRequest.Body).Decode(&requestData)
	if nil != err {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	author := model.NewAuthor(requestData.Name)
	if !author.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	add the author to database
	insertedAuthor, err := store.AddAuthor(author)
	if nil != err {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	fill response payload
	var responseData = authorResponse{}

	responseData.ID = insertedAuthor.Id
	responseData.Name = insertedAuthor.Name

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusCreated)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	get author by Id API
func GetAuthor(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)
	id := model.NewID(vars["id"])

	if !id.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	get the author by Id from database
	var author *model.Author

	author, err := store.GetAuthorByID(id)
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	var responseData = authorResponse{}

	responseData.ID = author.Id
	responseData.Name = author.Name

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	get all authors API
func GetAllAuthors(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	get all authors from database
	var authorList []model.Author

	authorList, err := store.GetAllAuthor()
	if err != nil {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	TODO: implement pagination

	//	fill response payload
	var responseData = authorListResponse{}

	for _, item := range authorList {

		var author = authorResponse{}

		author.ID = item.Id
		author.Name = item.Name

		responseData.Author = append(responseData.Author, author)
	}

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	patch author by Id API
func PatchAuthor(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)
	id := model.NewID(vars["id"])

	if !id.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	fetch request payload
	var requestData authorRequest

	err := json.NewDecoder(httpRequest.Body).Decode(&requestData)
	if nil != err {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(requestData.ID) > 0 && requestData.ID != string(*id) {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	author := model.NewAuthor(requestData.Name)
	if !author.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	update the author by Id in the database
	err = store.UpdateAuthor(id, author)
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	var responseData = authorResponse{}

	responseData.ID = string(*id)
	responseData.Name = author.Name

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	delete author by Id API
func DeleteAuthor(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)
	id := model.NewID(vars["id"])

	if !id.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	delete the author by Id from database
	err := store.DeleteAuthor(id)
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
}
