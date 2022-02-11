////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Feb/10/2022  -  aldebap
//
//	personLib entry point
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
)

////////////////////////////////////////////////////////////////////////////////
//	personLib entry point
////////////////////////////////////////////////////////////////////////////////

func main() {

	//	splash screen
	fmt.Printf(">>> personaLib Server\n\n")

	//	get configuration parameters from environment
	databaseURL := os.Getenv("DATABASEURL")
	servicePort := os.Getenv("SERVICEPORT")

	if len(servicePort) == 0 {
		servicePort = ":8080"
	} else if servicePort[0] != ':' {
		servicePort = ":" + servicePort
	}

	//	start personLib application
	personaLibApp := New(databaseURL)

	personaLibApp.Run(servicePort)
}
