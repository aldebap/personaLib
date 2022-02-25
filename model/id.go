////////////////////////////////////////////////////////////////////////////////
//	id.go  -  Feb/20/2022  -  aldebap
//
//	Object ID
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
func NewID(id string) *ID {
	var newId ID = ID(id)

	return &newId
}

//	validate object ID
func (id *ID) IsValid() bool {
	return idValidCharacters.MatchString(string(*id))
}
