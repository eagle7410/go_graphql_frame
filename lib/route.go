package lib

import (
	"fmt"
	"github.com/gorilla/mux"
	"go_graphql_frame/graphql"
	"net/http"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	//Graphql
	handlerGraphQl := graphql.GetHandlerGraphQl(&ENV)
	r.Handle("/graphql", handlerGraphQl)

	// Tech
	r.HandleFunc("/ping", ping)
	r.HandleFunc("/", toIndex)

	return r
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG \n  IP: %v\n  Host: %v\n", ReadUserIP(r), r.Host)
}

func toIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ping", http.StatusSeeOther)
}
