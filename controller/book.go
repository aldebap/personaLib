////////////////////////////////////////////////////////////////////////////////
//	book.go  -  Mar/11/2022  -  aldebap
//
//	personLib controller - book API
////////////////////////////////////////////////////////////////////////////////

package controller

import (
	"encoding/json"
	"net/http"
	"personaLib/model"
	"personaLib/store"

	"github.com/gorilla/mux"
)

//	book request
type bookRequest struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
}

//	book response
type bookResponse struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

//	book's list response
type bookListResponse struct {
	Book []bookResponse `json:"book"`
}

//	add book API
func AddBook(httpResponse http.ResponseWriter, httpRequest *http.Request) {

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
	var requestData bookRequest

	err := json.NewDecoder(httpRequest.Body).Decode(&requestData)
	if nil != err {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	book := model.NewBook(requestData.Title)
	if !book.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	add the book to database
	insertedBook, err := store.AddBook(book)
	if nil != err {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	fill response payload
	var responseData = bookResponse{}

	responseData.ID = insertedBook.Id
	responseData.Title = insertedBook.Title

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusCreated)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	get book by Id API
func GetBook(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)
	id := model.NewID(vars["id"])

	if !id.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	get the book by Id from database
	var book *model.Book

	book, err := store.GetBookByID(id)
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	var responseData = bookResponse{}

	responseData.ID = book.Id
	responseData.Title = book.Title

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	get all books API
func GetAllBooks(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	get all books from database
	var bookList []model.Book

	bookList, err := store.GetAllBook()
	if err != nil {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	TODO: implement pagination

	//	fill response payload
	var responseData = bookListResponse{}

	for _, item := range bookList {

		var book = bookResponse{}

		book.ID = item.Id
		book.Title = item.Title

		responseData.Book = append(responseData.Book, book)
	}

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	patch book by Id API
func PatchBook(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)
	id := model.NewID(vars["id"])

	if !id.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	fetch request payload
	var requestData bookRequest

	err := json.NewDecoder(httpRequest.Body).Decode(&requestData)
	if nil != err {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(requestData.ID) > 0 && requestData.ID != string(*id) {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	book := model.NewBook(requestData.Title)
	if !book.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	update the book by Id in the database
	err = store.UpdateBook(id, book)
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	var responseData = bookResponse{}

	responseData.ID = string(*id)
	responseData.Title = book.Title

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

//	delete book by Id API
func DeleteBook(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fetch request variables
	vars := mux.Vars(httpRequest)
	id := model.NewID(vars["id"])

	if !id.IsValid() {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}

	//	delete the book by Id from database
	err := store.DeleteBook(id)
	if err != nil {
		httpResponse.WriteHeader(http.StatusNotFound)
		return
	}

	//	fill response payload
	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
}
