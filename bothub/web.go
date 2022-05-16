package bothub

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var port = 80

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/stream-simulation", StreamSimulationHandler).Methods("GET")

	staticFileDirectory := http.Dir("./frontend/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.FileServer(staticFileDirectory)
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/").Handler(staticFileHandler)
	return r
}

func StartServer() {
	fmt.Println("Server started on port", port)

	r := NewRouter()
	http.ListenAndServe(":80", r)
}

func StreamSimulationHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}

	defer c.Close()
	RunSimulation(c, r)
}

func check(err error, name string) bool {
	if err != nil {
		fmt.Println(name, ": ", err.Error())
		return true
	} else {
		return false
	}
}
