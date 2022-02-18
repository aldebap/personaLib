////////////////////////////////////////////////////////////////////////////////
//	author.go  -  Feb/14/2022  -  aldebap
//
//	personLib controller - author API
////////////////////////////////////////////////////////////////////////////////

package controller

import (
	"encoding/json"
	"net/http"

	"personaLib/store"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	new author request
type newAuthorRequest struct {
	Name string `json:"name"`
}

//	author response
type authorResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//	fetch all authors response
type getAllAuthorResponse struct {
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
	var requestData newAuthorRequest

	err := json.NewDecoder(httpRequest.Body).Decode(&requestData)
	if nil != err {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	TODO: create a business rule to validate the payload before adding the author to store

	//	add the author to database
	var newAuthor store.Author

	newAuthor.Name = requestData.Name

	insertResult, err := store.AddAuthor(newAuthor)
	if nil != err {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	fill response payload
	var responseData = authorResponse{}

	responseData.ID = insertResult.Id.Hex()
	responseData.Name = insertResult.Name

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusCreated)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	get author by Id API
func GetAuthor(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)

	//	TODO: create a business rule to validate the id before query the author from store

	//	get the author by Id from database
	var author *store.Author

	author, err := store.GetAuthorByID(vars["id"])
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	var responseData = authorResponse{}

	responseData.ID = author.Id.Hex()
	responseData.Name = author.Name

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	get all authors API
func GetAllAuthors(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	get all authors from database
	var authorList []store.Author

	authorList, err := store.GetAllAuthor()
	if err != nil {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	TODO: implement pagination

	//	fill response payload
	var responseData = getAllAuthorResponse{}

	for _, item := range authorList {

		var authorData = authorResponse{}

		authorData.ID = item.Id.Hex()
		authorData.Name = item.Name

		responseData.Author = append(responseData.Author, authorData)
	}

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	patch author by Id API
func PatchAuthor(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)

	//	TODO: create a business rule to validate the id before query the author from store

	//	fetch request payload
	var requestData newAuthorRequest

	err := json.NewDecoder(httpRequest.Body).Decode(&requestData)
	if nil != err {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	TODO: create a business rule to validate the payload before adding the author to store

	//	TODO: if the Id is populated in the payload, it must be equal to the URL parameter

	//	update the author by Id in the database
	var author store.Author

	//	TODO: better to pass the Id as a string to updateAuthor() function
	author.Id, err = primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	author.Name = requestData.Name

	err = store.UpdateAuthor(author)
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	var responseData = authorResponse{}

	responseData.ID = author.Id.Hex()
	responseData.Name = author.Name

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	delete author by Id API
func DeleteAuthor(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)

	//	TODO: create a business rule to validate the id before query the author from store

	//	delete the author by Id from database
	err := store.DeleteAuthor(vars["id"])
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	var responseData = authorResponse{}

	responseData.ID = vars["id"]

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}
