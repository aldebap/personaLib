////////////////////////////////////////////////////////////////////////////////
//	author.go  -  Feb/18/2022  -  aldebap
//
//	Author entity
////////////////////////////////////////////////////////////////////////////////

package model

import (
	"regexp"
)

var (
	authorsValidCharacters *regexp.Regexp
)

//	author attributes
type Author struct {
	Id   string
	Name string
}

func init() {
	//	TODO: add portuguese accent chars to regexp
	authorsValidCharacters = regexp.MustCompile("^[a-zA-Z 0-9.]+$")
}

//	create a new Author
func NewAuthor(name string) *Author {
	return &Author{
		Name: name,
	}
}

//	validate author's fields
func (author *Author) IsValid() bool {
	//	name field validation
	if len(author.Name) == 0 {
		return false
	}
	//	TODO: move the max size to the regexp
	if len(author.Name) > 50 {
		return false
	}
	if !authorsValidCharacters.MatchString(author.Name) {
		return false
	}

	return true
}
