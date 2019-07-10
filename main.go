package main

import (
	"github.com/gorilla/handlers"
	"go_frame/lib"
	"log"
	"net/http"
)

const port = ":8080"

func init() {

	lib.OpenLogFile()

	if err := lib.ENV.Init(); err != nil {
		lib.LogFatalf("Error on initializing environment : %s", err)
	}
}

func main() {

	router := lib.GetRouter()

	lib.LogAppRun(port)

	middleware := lib.SetCorsMiddleware(
		lib.LogRequest(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"*"}),
				handlers.AllowedMethods(lib.AllowedMethods()),
				handlers.AllowedOrigins([]string{"*"}))(router)))

	log.Fatal(http.ListenAndServe(port, middleware))
}
