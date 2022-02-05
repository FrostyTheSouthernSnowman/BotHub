package robot

import (
	"fmt"
	"math"
	"strings"
)

const (
	//NORTH constant to be used for facing
	NORTH = "NORTH"
	//SOUTH constant to be used for facing
	SOUTH = "SOUTH"
	//EAST constant to be used for facing
	EAST = "EAST"
	//WEST constant to be used for facing
	WEST = "WEST"

	errRobotNotPlaced  = "Robot is not placed on the table yet."
	errRobotOverBoard  = "Command ignored, Robot will fall."
	errMovementIgnored = "Invalid Movement so it is ignored."
	errInvalidCmd      = "Sorry i don't get that!"
	errFailToInitiate  = "Failed to initiate"
)

//Robot hold all the information about the robot
type Robot struct {
	X, Y          int
	F             string
	IsRobotPlaced bool
	Table         Table
}

//Perform is to receive a command and perform the command
func (r *Robot) Perform(command string, x int, y int, f string) (int, int, string, error) {
	var err error
	var rbot_x int
	var rbot_y int
	var rbot_f string
	switch command {
	case "PLACE":
		rbot_x, rbot_y, rbot_f, err = r.Place(x, y, f)
	case "MOVE":
		rbot_x, rbot_y, rbot_f, err = r.Move()
	case "LEFT":
		rbot_x, rbot_y, rbot_f, err = r.Left()
	case "RIGHT":
		rbot_x, rbot_y, rbot_f, err = r.Right()
	case "REPORT":
		return r.Report()
	default:
		return 0, 0, "", fmt.Errorf(errInvalidCmd)
	}
	if err != nil {
		return 0, 0, "", err
	}
	return rbot_x, rbot_y, rbot_f, nil
}

//Place will put the toy robot on the table in position X,Y and facing NORTH, SOUTH, EAST or WEST.
func (r *Robot) Place(x int, y int, f string) (int, int, string, error) {
	direction := strings.ToUpper(f)
	//IF its not on the table
	if !isStillOnTheTable(x, y, r.Table) {
		return 0, 0, "", fmt.Errorf(errRobotOverBoard)
	}
	if !isValidFacing(direction) {
		return 0, 0, "", fmt.Errorf(errInvalidCmd)
	}
	r.X = x
	r.Y = y
	r.F = direction
	r.IsRobotPlaced = true
	return r.X, r.Y, r.F, nil
}

//Move will move the toy robot one unit forward in the direction it is currently facing.
func (r *Robot) Move() (int, int, string, error) {
	if !r.IsRobotPlaced {
		return 0, 0, "", fmt.Errorf(errRobotNotPlaced)
	}
	switch r.F {
	case NORTH:
		//IF after move and still in the table
		if isStillOnTheTable(r.X, r.Y+1, r.Table) {
			r.Y++
			return r.X, r.Y, r.F, nil
		}
	case SOUTH:
		//IF after move and still in the table
		if isStillOnTheTable(r.X, r.Y-1, r.Table) {
			r.Y--
			return r.X, r.Y, r.F, nil
		}
	case EAST:
		//IF after move and still in the table
		if isStillOnTheTable(r.X+1, r.Y, r.Table) {
			r.X++
			return r.X, r.Y, r.F, nil
		}
	case WEST:
		//IF after move and still in the table
		if isStillOnTheTable(r.X-1, r.Y, r.Table) {
			r.X--
			return r.X, r.Y, r.F, nil
		}
	}
	return 0, 0, "", fmt.Errorf(errMovementIgnored)
}

//Left will rotate the robot 90 degrees to left in the specified direction without changing the position of the robot.
func (r *Robot) Left() (int, int, string, error) {
	if !r.IsRobotPlaced {
		return 0, 0, "", fmt.Errorf(errRobotNotPlaced)
	}
	switch r.F {
	case NORTH:
		r.F = WEST
	case SOUTH:
		r.F = EAST
	case EAST:
		r.F = NORTH
	case WEST:
		r.F = SOUTH
	}
	return r.X, r.Y, r.F, nil
}

//Right will rotate the robot 90 degrees to right in the specified direction without changing the position of the robot.
func (r *Robot) Right() (int, int, string, error) {
	if !r.IsRobotPlaced {
		return 0, 0, "", fmt.Errorf(errRobotNotPlaced)
	}
	switch r.F {
	case NORTH:
		r.F = EAST
	case SOUTH:
		r.F = WEST
	case EAST:
		r.F = SOUTH
	case WEST:
		r.F = NORTH
	}
	return r.X, r.Y, r.F, nil
}

//Report will announce the X,Y and F of the robot. This can be in any form, but standard output is sufficient
func (r *Robot) Report() (int, int, string, error) {
	if !r.IsRobotPlaced {
		return 0, 0, "", fmt.Errorf(errRobotNotPlaced)
	}
	return r.X, r.Y, r.F, nil
}

//NewRobot is to instantiate a new robot object
func NewRobot(tableWidth int, tableLength int) (Robot, error) {
	table, err := NewTable(tableWidth, tableLength)
	if err != nil {
		return Robot{}, fmt.Errorf(errFailToInitiate)
	}
	return Robot{Table: table}, nil
}

//private methods

func isStillOnTheTable(x int, y int, table Table) bool {
	return int(math.Abs(float64(x))) < table.Width/2 && int(math.Abs(float64(y))) < table.Length/2
}

func isValidFacing(f string) bool {
	return f == NORTH || f == SOUTH || f == EAST || f == WEST
}
