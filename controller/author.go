////////////////////////////////////////////////////////////////////////////////
//	author.go  -  Feb/14/2022  -  aldebap
//
//	personLib controller - author API
////////////////////////////////////////////////////////////////////////////////

package controller

import (
	"encoding/json"
	"net/http"

	"personaLib/entity"

	"github.com/gorilla/mux"
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

	//	add the author to database
	var newAuthor entity.Author

	newAuthor.Name = requestData.Name

	insertResult, err := entity.AddAuthor(newAuthor)
	if nil != err {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	fill response payload
	var responseData = authorResponse{}

	responseData.ID = insertResult.Id
	responseData.Name = insertResult.Name

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusCreated)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	get author by Id API
func GetAuthor(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)

	//	get all authors from database
	var author *entity.Author

	author, err := entity.GetAuthorByID(vars["id"])
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
	var authorList []entity.Author

	authorList, err := entity.GetAllAuthor()
	if err != nil {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	fill response payload
	var responseData = getAllAuthorResponse{}

	for _, item := range authorList {

		var authorData = authorResponse{}

		authorData.ID = item.Id
		authorData.Name = item.Name

		responseData.Author = append(responseData.Author, authorData)
	}

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}
