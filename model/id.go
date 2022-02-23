////////////////////////////////////////////////////////////////////////////////
//	id.go  -  Feb/20/2022  -  aldebap
//
//	Entity ID
////////////////////////////////////////////////////////////////////////////////

package model

import "regexp"

var (
	idValidCharacters *regexp.Regexp
)

//	entity ID
type ID string

//	compile the validation regexp
func init() {
	idValidCharacters = regexp.MustCompile("^[a-fA-F0-9]{24}$")
}

//	create a new model ID from string
func FromString(id string) *ID {
	var newId ID = ID(id)

	return &newId
}

//	validate entity ID
func (id *ID) IsValid() bool {
	return idValidCharacters.MatchString(string(*id))
}
