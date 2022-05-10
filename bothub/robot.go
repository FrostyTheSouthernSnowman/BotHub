package bothub

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const num_objects int = 1

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type Message struct {
	Type string `json:"type"`
}

var messages [1]Message

type RigidBodySphereBoundingBox struct {
	Position Vector3
	Velocity Vector3
	Mass     float32
	Radius   float32
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

// Global array of particles.
var simulation_objects [num_objects]RigidBodySphereBoundingBox

const floor_height = 0

func ThreeDimensionalEuclidianDistance(point1 Vector3, point2 Vector3) float32 {
	return float32(math.Sqrt(float64(math.Pow(float64(point1.X-point2.X), 2) + math.Pow(float64(point1.Y-point2.Y), 2) + math.Pow(float64(point1.Z-point2.Z), 2))))
}

func PrintObjects() {
	for i := 0; i < num_objects; i++ {
		var object *RigidBodySphereBoundingBox = &simulation_objects[i]
		fmt.Printf("object[%v] (%.2f, %.2f, %.2f)\n", i, object.Position.X, object.Position.Y, object.Position.Z)
	}
}

// Initialize all objects with random positions and velocities, 1kg mass, and radius 1
//! This method will be removed in the near future
// TODO: Rewrite this function to initialize objects that aren't sphere and change the defaults to something specified by the user
func InitializeObjects() {
	for i := 0; i < num_objects; i++ {
		simulation_objects[i].Position = Vector3{float32(rand.Intn(50)), float32(rand.Intn(50)), float32(rand.Intn(50))}
		simulation_objects[i].Velocity = Vector3{float32(rand.Intn(20)), float32(rand.Intn(20)), 0}
		simulation_objects[i].Mass = 1
		simulation_objects[i].Radius = 1
	}
}

// Just applies Earth's gravity force (mass times gravity acceleration 9.81 m/s^2) to each particle.
func AddGravity(object *RigidBodySphereBoundingBox) Vector3 {
	return Vector3{0, 0, object.Mass * -9.81}
}

func CheckIfCollisionOccurred(object1 *RigidBodySphereBoundingBox, object2 *RigidBodySphereBoundingBox) bool {
	if ThreeDimensionalEuclidianDistance(object1.Position, object2.Position) < object1.Radius+object2.Radius {
		return true
	}
	return false
}

func RunSimulation(c *websocket.Conn, r *http.Request) {
	var totalSimulationTime float32 = 60 // The simulation will run for 10 seconds.
	var currentTime float32 = 0          // This accumulates the time that has passed.
	var dt float32 = 0.03333333333333333 // FPS raised to -1

	InitializeObjects()
	initial_simulation_state := simulation_objects
	PrintObjects()

	go ReadDataGorutine(c)

	for currentTime < totalSimulationTime {
		// We're sleeping here to keep things simple. In real applications you'd use some
		// timing API to get the current time in milliseconds and compute dt in the beginning
		// of every iteration like this:
		// currentTime = GetTime()
		// dt = currentTime - previousTime
		// previousTime = currentTime
		time.Sleep(time.Duration(dt*1000) * time.Millisecond)

		if messages[0].Type == "pause" {
			for {
				if messages[0].Type == "play" {
					break
				} else if messages[0].Type == "reset" {
					simulation_objects = initial_simulation_state
					json_formatted_sim, err := json.Marshal(simulation_objects)

					err = c.WriteMessage(websocket.TextMessage, json_formatted_sim)
					if err != nil {
						fmt.Println("error:", err)
						return
					}
				}
			}
		}

		for i := 0; i < num_objects; i++ {
			var object *RigidBodySphereBoundingBox = &simulation_objects[i]
			var force Vector3 = AddGravity(object)
			var acceleration Vector3 = Vector3{force.X / object.Mass, force.Y / object.Mass, force.Z / object.Mass}
			object.Velocity.X += acceleration.X * dt
			object.Velocity.Y += acceleration.Y * dt
			object.Velocity.Z += acceleration.Z * dt
			object.Position.X += object.Velocity.X * dt
			object.Position.Y += object.Velocity.Y * dt
			object.Position.Z += object.Velocity.Z * dt

			if object.Position.Z-object.Radius <= floor_height {
				object.Velocity.Z = object.Velocity.Z * -1
				object.Position.Z = 0 + object.Radius
			}

			//! Walls will temporarily be placed at 50 x and 50 y
			if object.Position.X+object.Radius >= 50 {
				object.Velocity.X = object.Velocity.X * -1
			}

			if object.Position.Y+object.Radius >= 50 {
				object.Velocity.Y = object.Velocity.Y * -1
			}

			if object.Position.X-object.Radius <= 0 {
				object.Velocity.X = object.Velocity.X * -1
			}

			if object.Position.Y-object.Radius <= 0 {
				object.Velocity.Y = object.Velocity.Y * -1
			}

			PrintObjects()
			currentTime += dt
		}

		for i := 0; i < num_objects; i++ {

			var object1 *RigidBodySphereBoundingBox = &simulation_objects[i]

			for i2 := 0; i2 < num_objects; i2++ {
				if i2 == i {
					continue
				}

				var object2 *RigidBodySphereBoundingBox = &simulation_objects[i2]
				if CheckIfCollisionOccurred(object1, object2) {
					fmt.Println("A collission occurred!")
					fmt.Println("Distance:", ThreeDimensionalEuclidianDistance(object1.Position, object2.Position))
					fmt.Println("Object 1 position:", object1.Position)
					fmt.Println("Object 2 position:", object2.Position)
				}
			}
		}

		json_formatted_sim, err := json.Marshal(simulation_objects)

		err = c.WriteMessage(websocket.TextMessage, json_formatted_sim)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	}
}
