////////////////////////////////////////////////////////////////////////////////
//	controller.go  -  Feb/11/2022  -  aldebap
//
//	personLib routes controller
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"net/http"

	"personaLib/author"
)

//	get author response
type getAccountResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//	fetch all authors response
type getAllAuthorResponse struct {
	Author []getAccountResponse `json:"author"`
}

//	get all authors API
func getAllAccounts(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	get all authors from database
	var authorList []author.Author

	authorList, err := author.GetAll()
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	var responseData = getAllAuthorResponse{}

	for _, item := range authorList {

		var authorData = getAccountResponse{}

		authorData.ID = item.Id
		authorData.Name = item.Name

		responseData.Author = append(responseData.Author, authorData)
	}

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}
