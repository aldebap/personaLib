////////////////////////////////////////////////////////////////////////////////
//	publisher.go  -  Mar/01/2022  -  aldebap
//
//	Publisher entity
////////////////////////////////////////////////////////////////////////////////

package model

import (
	"regexp"
)

var (
	publishersValidCharacters *regexp.Regexp
)

//	publisher attributes
type Publisher struct {
	Id   string
	Name string
}

//	compile the validation regexp
func init() {
	publishersValidCharacters = regexp.MustCompile("^[a-zA-ZàáãéêíóôúüçÀÁÃÉÊÍÓÔÚÜÇ 0-9.]{1,50}$")
}

//	create a new Publisher
func NewPublisher(name string) *Publisher {
	return &Publisher{
		Name: name,
	}
}

//	validate publisher's fields
func (publisher *Publisher) IsValid() bool {
	//	name field validation
	if len(publisher.Name) == 0 {
		return false
	}
	if !publishersValidCharacters.MatchString(publisher.Name) {
		return false
	}

	return true
}
