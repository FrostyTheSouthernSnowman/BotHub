package bothub

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var num_objects int = 1

var addObject bool = false

type Message struct {
	Type string `json:"type"`
}

var messages [1]Message

type RigidBodySphereBoundingBox struct {
	Position Vector3
	Velocity Vector3
	Mass     float64
	Radius   float64
}

func ReadDataGorutine(c *websocket.Conn) {
	for {
		var message Message
		err := c.ReadJSON(&message)
		if err != nil {
			break
		}

		if message.Type == "addObject" {
			addObject = true
		} else {
			messages[0] = message
		}
	}
}

// Global array of particles.
var simulation_objects []RigidBodySphereBoundingBox

const floor_height = 0

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
	num_objects = 1
	simulation_objects = nil
	for i := 0; i < num_objects; i++ {
		var new_object = RigidBodySphereBoundingBox{
			Position: Vector3{float64(rand.Intn(50)), float64(rand.Intn(50)), float64(rand.Intn(50))},
			Velocity: Vector3{float64(rand.Intn(20)), float64(rand.Intn(20)), 0},
			Mass:     1,
			Radius:   1,
		}
		simulation_objects = append(simulation_objects, new_object)
	}
	addObject = false
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

func AddObject(initial_state []RigidBodySphereBoundingBox, c *websocket.Conn) []RigidBodySphereBoundingBox {
	new_object := RigidBodySphereBoundingBox{
		Position: Vector3{
			X: float64(rand.Intn(50)),
			Y: float64(rand.Intn(50)),
			Z: float64(rand.Intn(50)),
		},
		Velocity: Vector3{
			X: float64(rand.Intn(50)),
			Y: float64(rand.Intn(50)),
			Z: 0,
		},
		Mass:   1,
		Radius: 1,
	}

	simulation_objects = append(simulation_objects, new_object)
	initial_state = append(initial_state, new_object)
	num_objects++

	json_formatted_sim, err := json.Marshal(simulation_objects)

	check(err, "could not marshall simulation_objects to JSON when trying to add a new object")

	err = c.WriteMessage(websocket.TextMessage, json_formatted_sim)
	if err != nil {
		fmt.Println("error:", err)
		return initial_state
	}
	addObject = false
	return initial_state
}

func RunSimulation(c *websocket.Conn, r *http.Request) {
	var dt float64 = 0.03333333333333333 // FPS raised to -1
	messages[0] = Message{
		Type: "pause",
	}

	var points []*Point
	var sticks []*Stick

	a := MakePoint(Vector3{20, 10, 100}, 0.9, 0.99)
	b := MakePoint(Vector3{30, 10, 100}, 0.9, 0.99)
	point_c := MakePoint(Vector3{30, 20, 100}, 0.9, 0.99)
	d := MakePoint(Vector3{20, 20, 100}, 0.9, 0.99)

	e := MakePoint(Vector3{20, 10, 110}, 0.9, 0.99)
	f := MakePoint(Vector3{30, 10, 110}, 0.9, 0.99)
	g := MakePoint(Vector3{30, 20, 110}, 0.9, 0.99)
	h := MakePoint(Vector3{20, 20, 110}, 0.9, 0.99)

	points = append(points, &a)
	points = append(points, &b)
	points = append(points, &point_c)
	points = append(points, &d)

	points = append(points, &e)
	points = append(points, &f)
	points = append(points, &g)
	points = append(points, &h)

	stick01 := MakeStick(&a, &b)
	stick02 := MakeStick(&b, &point_c)
	stick03 := MakeStick(&point_c, &d)
	stick04 := MakeStick(&d, &a)

	stick05 := MakeStick(&e, &f)
	stick06 := MakeStick(&f, &g)
	stick07 := MakeStick(&g, &h)
	stick08 := MakeStick(&h, &e)

	stick09 := MakeStick(&a, &e)
	stick10 := MakeStick(&b, &f)
	stick11 := MakeStick(&point_c, &g)
	stick12 := MakeStick(&d, &h)

	stick13 := MakeStick(&a, &g)
	stick14 := MakeStick(&b, &h)
	stick15 := MakeStick(&point_c, &e)
	stick16 := MakeStick(&d, &f)

	stick17 := MakeStick(&a, &point_c)
	stick18 := MakeStick(&b, &g)
	stick19 := MakeStick(&f, &h)
	stick20 := MakeStick(&d, &e)

	stick21 := MakeStick(&d, &g)
	stick22 := MakeStick(&a, &f)

	sticks = append(sticks, &stick01)
	sticks = append(sticks, &stick02)
	sticks = append(sticks, &stick03)
	sticks = append(sticks, &stick04)

	sticks = append(sticks, &stick05)
	sticks = append(sticks, &stick06)
	sticks = append(sticks, &stick07)
	sticks = append(sticks, &stick08)

	sticks = append(sticks, &stick09)
	sticks = append(sticks, &stick10)
	sticks = append(sticks, &stick11)
	sticks = append(sticks, &stick12)

	sticks = append(sticks, &stick13)
	sticks = append(sticks, &stick14)
	sticks = append(sticks, &stick15)
	sticks = append(sticks, &stick16)

	sticks = append(sticks, &stick17)
	sticks = append(sticks, &stick18)
	sticks = append(sticks, &stick19)
	sticks = append(sticks, &stick20)

	sticks = append(sticks, &stick21)
	sticks = append(sticks, &stick22)

	var max_time float64 = 30
	var current_time float64

	for current_time < max_time {
		// We're sleeping here to keep things simple. In real applications you'd use some
		// timing API to get the current time in milliseconds and compute dt in the beginning
		// of every iteration like this:
		// currentTime = GetTime()
		// dt = currentTime - previousTime
		// previousTime = currentTime
		time.Sleep(time.Duration(dt*1000) * time.Millisecond)
		current_time += dt

		for _, point := range points {
			point.Update(dt)
		}

		for _, stick := range sticks {
			stick.Update()
		}

		json_formatted_sim, err := json.Marshal(sticks)

		err = c.WriteMessage(websocket.TextMessage, json_formatted_sim)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	}
}
