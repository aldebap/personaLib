////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Feb/10/2022  -  aldebap
//
//	personLib entry point
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
//	personLib entry point
////////////////////////////////////////////////////////////////////////////////

func main() {

	//	splash screen
	fmt.Printf(">>> personaLib Server\n\n")

	//	start personLib application
	var config = GetFromEnv()
	var personaLibApp App

	personaLibApp.Run(config)
}
