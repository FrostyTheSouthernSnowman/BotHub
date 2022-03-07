package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"robot-simulator/robot"
	"robot-simulator/robot/physics"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

var sim_robots []robot.Robot

var messages []Message = []Message{{}}

type Command struct {
	Commamd string `json:"command"`
}

type Message struct {
	Type string `json:"type"`
}

type XYPosition struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	F string  `json:"f"`
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/set-position", SetPositionHandler).Methods("POST")
	r.HandleFunc("/api/place-robot", PlaceRobotHandler).Methods("POST")
	r.HandleFunc("/api/stream-simulation", StreamHandler).Methods("GET")
	r.HandleFunc("/api/add-robot", AddRobotHandler).Methods("POST")

	staticFileDirectory := http.Dir("./frontend/built")
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

	fmt.Println("Starting server")
	r := NewRouter()
	fmt.Println("Server started on port 80")
	http.ListenAndServe(":80", r)
}

func SetPositionHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var x_and_y XYPosition
	err := decoder.Decode(&x_and_y)
	check(err, "decoding in SetPositionHandler")
	//initialize table & robot
	rbot, err := robot.NewRobot(x_and_y.X, x_and_y.Y)
	if err != nil {
		panic(err)
	}

	sim_robots = []robot.Robot{rbot}

	fmt.Fprintf(w, "{\"errors\": \"false\"}")
}

func PlaceRobotHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var robot_pos XYPosition
	err := decoder.Decode(&robot_pos)
	check(err, "decode json in PlaceRobotHandler")
	robot_pos.X, robot_pos.Y, robot_pos.F, err = sim_robots[0].Perform("PLACE", robot_pos.X, robot_pos.Y, 5.00, robot_pos.F)
	if !check(err, "can't place robot") {
		robot_bytes, err := json.Marshal(robot_pos)
		if !check(err, "error marshalling bytes") {
			w.Write(robot_bytes)
		}
	}
}

func AddRobotHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var new_robot_pos physics.Vector3
	err := decoder.Decode(&new_robot_pos)
	check(err, "decode robot position json in AddRobotHandler")
	rbot := robot.Robot{
		Position: physics.XYZPosition{
			X:         new_robot_pos.X,
			Y:         new_robot_pos.Y,
			Z:         new_robot_pos.Z,
			XRotation: 0,
			YRotation: 0,
			ZRotation: 0,
			Velocity: physics.Vector3{
				X: 0,
				Y: 0,
				Z: 0,
			},
		},
		X:             int(new_robot_pos.X),
		Y:             int(new_robot_pos.Y),
		F:             "EAST", //! This is temporary, eventually f will be deprecated because it serves no purpose. f should be replaced by a quaternion for orientation.
		IsRobotPlaced: true,
		Table:         sim_robots[0].Table,
	}
	sim_robots = append(sim_robots, rbot)
	new_robot_pos.X, new_robot_pos.Y, _, err = sim_robots[len(sim_robots)-1].Perform("PLACE", new_robot_pos.X, new_robot_pos.Y, new_robot_pos.Z, "EAST")
	if !check(err, "can't place robot") {
		robot_bytes, err := json.Marshal(new_robot_pos)
		if !check(err, "error marshalling bytes") {
			w.Write(robot_bytes)
		}
	}
}

func ReadDataGorutine(c *websocket.Conn) {
	for {
		var message Message
		err := c.ReadJSON(&message)
		if err != nil {
			break
		}
		messages[0] = message
	}
}

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	defer c.Close()

	sim_robots[0].Position = physics.XYZPosition{
		X:         sim_robots[0].Position.X,
		Y:         sim_robots[0].Position.Y,
		Z:         1.00,
		XRotation: 0,
		YRotation: 0,
		ZRotation: 0,
		Velocity: physics.Vector3{
			X: 0,
			Y: 0,
			Z: 0,
		},
	}
	var initial_state []robot.Robot = sim_robots

	go ReadDataGorutine(c)

	for {
		// Determine if pause, play, or reset is necessary
		if messages[0].Type != "" {
			switch messages[0].Type {
			case "pause":
				time.Sleep(time.Duration(time.Duration(time.Second / time.Duration((physics.CalculationsPerSecond)))))
				continue

			case "reset":
				sim_robots[0].Position = initial_state[0].Position
				messages[0].Type = "pause"

				sim_json, err := json.Marshal([]physics.XYZPosition{initial_state[0].Position})
				check(err, "marshall json")

				// Main loop, send new data to the frontend
				err = c.WriteMessage(websocket.TextMessage, sim_json)
				if err != nil {
					fmt.Println("error:", err)
					break
				}
				time.Sleep(time.Duration(time.Duration(time.Second / time.Duration((physics.CalculationsPerSecond)))))
				continue

			case "play":
				messages[0].Type = ""
				time.Sleep(time.Duration(time.Duration(time.Second / time.Duration((physics.CalculationsPerSecond)))))
				continue
			}
		}

		sim, _ := physics.CalculatePhysics([]physics.XYZPosition{sim_robots[0].Position})
		sim_robots[0].Position = sim[0]

		sim_json, err := json.Marshal(sim)
		check(err, "marshall json")

		// Send new data to the frontend
		err = c.WriteMessage(websocket.TextMessage, sim_json)
		if err != nil {
			fmt.Println("error:", err)
			break
		}
		time.Sleep(time.Duration(time.Duration(time.Second / time.Duration((physics.CalculationsPerSecond)))))
	}
}

func check(err error, name string) bool {
	if err != nil {
		fmt.Println(name, ": ", err.Error())
		return true
	} else {
		return false
	}
}
