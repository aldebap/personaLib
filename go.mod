module personaLib

go 1.17

require (
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/klauspost/compress v1.14.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.0 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.mongodb.org/mongo-driver v1.8.3 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/text v0.3.7 // indirect
)
require personaLib/controller v0.0.0-unpublished
replace personaLib/controller v0.0.0-unpublished => ./controller
require personaLib/model v0.0.0-unpublished
replace personaLib/model v0.0.0-unpublished => ./model
require personaLib/store v0.0.0-unpublished
replace personaLib/store v0.0.0-unpublished => ./store
