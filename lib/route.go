
package lib

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	//TODO: clear Do somethining ...
	//r.PathPrefix("/static/").Handler(
	//	staticAccess(
	//		http.StripPrefix(
	//			"/static/",
	//			http.FileServer(http.Dir(path.Join(ENV.WorkDir, "/front/dist"))))))

	// Tech
	r.HandleFunc("/ping", ping)
	r.HandleFunc("/", toIndex)

	return r
}

//TODO: clear
//func staticAccess(handler http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		//TODO: clear Maybe check access
//		handler.ServeHTTP(w, r)
//
//	})
//}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG \n  IP: %v\n  Host: %v\n", ReadUserIP(r), r.Host)
}

func toIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ping", http.StatusSeeOther)
}
