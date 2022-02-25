////////////////////////////////////////////////////////////////////////////////
//	author_test.go  -  Feb/20/2022  -  aldebap
//
//	Author entity tests
////////////////////////////////////////////////////////////////////////////////

package model

import (
	"fmt"
	"testing"
)

//	scenario 01 - test name validation
func TestNameValidation(t *testing.T) {
	t.Run(">>> author's name validation", func(t *testing.T) {

		//	some test cases
		var testScenarios = []struct {
			input  string
			result bool
		}{
			{input: "", result: false},
			{input: "123456789012345678901234567890123456789012345678901", result: false},
			{input: "Joao da Silva & Souza", result: false},
			{input: "Joao da Silva ' or true '", result: false},
			{input: "12345678901234567890123456789012345678901234567890", result: true},
			{input: "Joao da Silva", result: true},
			{input: "João José da Silva", result: true},
		}

		for i, test := range testScenarios {
			fmt.Printf("scenario: 01.%02d name: '%s'\texpected: %t\n", i+1, test.input, test.result)

			author := NewAuthor(test.input)

			want := test.result
			got := author.IsValid()

			if want != got {
				t.Errorf("Author's validation: got %t, want %t", got, want)
			}
		}
	})
}

//	scenario 02 - test to check that a buffer overflow attempt is not valid
func TestBufferOverflow(t *testing.T) {
	t.Run(">>> buffer overflow on name", func(t *testing.T) {

		author := NewAuthor("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

		want := false
		got := author.IsValid()
		if want != got {
			t.Errorf("Author's validation: got %t, want %t", got, want)
		}
	})
}