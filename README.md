# base64api

About:
A small example of a Rest API developed in Go Lang where it encodes and decodes 64 base text.

Dependency:
	"github.com/gorilla/mux"


build-linux:
	env GOOS=linux go build -o neon-auth-api ./cmd/api/main.go

build:
	go build -o base64api ./cmd/api/main.go


Usage:

To encode:
With an application running on localhost for example, send a request with the POST method to http://localhost:8181/encode
containing Json in the body with the content to be coded:

{
	"encodeTXT": "text example"
}


To decode:
Send a request with the POST method to http://localhost:8181/decode
containing Json in the body with the content to be decoded:

{
	"decodeTXT": "dGV4dCBleGFtcGxl"
}

To Generate password:
Send a request with the GET method to http://localhost:8181/pwd-random/{pwd_size}
pwd_size is the password length to generate:
ex: http://localhost:8181/pwd-random/10
