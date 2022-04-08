////////////////////////////////////////////////////////////////////////////////
//	book.go  -  Mar/14/2022  -  aldebap
//
//	Book entity
////////////////////////////////////////////////////////////////////////////////

package model

import (
	"regexp"
)

var (
	bookValidCharacters *regexp.Regexp
)

//	book attributes
type Book struct {
	Id    string
	Title string
}

//	compile the validation regexp
func init() {
	bookValidCharacters = regexp.MustCompile("^[a-zA-ZàáãéêíóôúüçÀÁÃÉÊÍÓÔÚÜÇ 0-9.:-]{1,50}$")
}

//	create a new Book
func NewBook(title string) *Book {
	return &Book{
		Title: title,
	}
}

//	validate book's fields
func (book *Book) IsValid() bool {
	//	title field validation
	if len(book.Title) == 0 {
		return false
	}
	if !bookValidCharacters.MatchString(book.Title) {
		return false
	}

	return true
}
