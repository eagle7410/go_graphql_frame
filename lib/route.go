package lib

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
	"go_graphql_frame/graphql"
	"net/http"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.Handle("/graphql", handler.New(&handler.Config{
		Schema:     graphql.Schema.GetSchemaLink(),
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	}))

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
