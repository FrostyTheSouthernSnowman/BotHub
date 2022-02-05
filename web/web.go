package web

import (
	// "fmt" has methods for formatted I/O operations (like printing to the console)

	// The "net/http" library has methods to implement HTTP clients and servers

	"encoding/json"
	"fmt"
	"net/http"

	"robot-simulator/robot"

	"github.com/gorilla/mux"
)

var sim_robot robot.Robot

type Command struct {
	Commamd string `json:"command"`
}

type XYPosition struct {
	X int    `json:"x"`
	Y int    `json:"y"`
	F string `json:"f"`
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/set-position", SetPositionHandler).Methods("POST")
	r.HandleFunc("/api/move-robot", MoveRobotHandler).Methods("POST")
	r.HandleFunc("/api/place-robot", PlaceRobotHandler).Methods("POST")

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
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the c
	//Handle all the available flags
	//fileName := flag.String("f", "", "A file name to be executed.")

	r := NewRouter()
	http.ListenAndServe(":80", r)
}

func SetPositionHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var x_and_y XYPosition
	err := decoder.Decode(&x_and_y)
	//initialize table & robot
	rbot, err := robot.NewRobot(x_and_y.X, x_and_y.Y)
	if err != nil {
		panic(err)
	}

	sim_robot = rbot
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}

func MoveRobotHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var robot_pos XYPosition
	var command Command
	err := decoder.Decode(&command)
	check(err)
	robot_pos.X, robot_pos.Y, robot_pos.F, err = sim_robot.Perform(command.Commamd, 0, 0, "")
	if !check(err) {
		robot_bytes, err := json.Marshal(robot_pos)
		if !check(err) {
			w.Write(robot_bytes)
		}
	}
}

func PlaceRobotHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var robot_pos XYPosition
	err := decoder.Decode(&robot_pos)
	check(err)
	robot_pos.X, robot_pos.Y, robot_pos.F, err = sim_robot.Perform("PLACE", robot_pos.X, robot_pos.Y, robot_pos.F)
	if !check(err) {
		robot_bytes, err := json.Marshal(robot_pos)
		if !check(err) {
			w.Write(robot_bytes)
		}
	}
}

func check(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}
