package lib

import (
	util "github.com/eagle7410/go_util/libs"
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
	r.HandleFunc("/ping", util.Ping)
	r.HandleFunc("/", toIndex)

	return r
}

func toIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ping", http.StatusSeeOther)
}
