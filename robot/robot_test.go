package robot_test

import (
	. "robot-simulator/robot"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Robot", func() {
	Describe("Invalid Command", func() {
		robot := initialiseRobot()
		_, _, _, err := robot.Perform("JUMP", 0, 0, "")
		It("should return err", func() {
			Expect(err != nil).To(Equal(true))
		})
	})

	Describe("Place Command", func() {
		robot := initialiseRobot()
		Context("with Invalid Direction", func() {
			_, _, _, err := robot.Perform("PLACE", 0, 0, "CENTER")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})
		Context("with X over the table", func() {
			_, _, _, err := robot.Perform("PLACE", 5, 0, "NORTH")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Context("with Y over the table", func() {
			_, _, _, err := robot.Perform("PLACE", 0, 5, "NORTH")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Context("with valid param", func() {
			x, y, f, err := robot.Perform("PLACE", 1, 2, "NORTH")
			It("should return no error", func() {
				Expect(err == nil).To(Equal(true))
				Expect(x == 1).To(Equal(true))
				Expect(y == 2).To(Equal(true))
				Expect(f == "NORTH").To(Equal(true))
			})
		})
	})

	Describe("Move Command", func() {
		Context("without placing the robot on the table", func() {
			robot := initialiseRobot()
			_, _, _, err := robot.Perform("MOVE", 0, 0, "")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Describe("with invalid move", func() {
			Context("when facing north", func() {
				robot := placeRobot(0, 4, "NORTH")
				_, _, _, err := robot.Perform("MOVE", 0, 0, "")
				It("should return err", func() {
					Expect(err != nil).To(Equal(true))
				})
			})
			Context("when facing south", func() {
				robot := placeRobot(0, -4, "SOUTH")
				_, _, _, err := robot.Perform("MOVE", 0, 0, "")
				It("should return err", func() {
					Expect(err != nil).To(Equal(true))
				})
			})
			Context("when facing east", func() {
				robot := placeRobot(4, 0, "EAST")
				_, _, _, err := robot.Perform("MOVE", 0, 0, "")
				It("should return err", func() {
					Expect(err != nil).To(Equal(true))
				})
			})
			Context("when facing west", func() {
				robot := placeRobot(-4, 0, "WEST")
				_, _, _, err := robot.Perform("MOVE", 0, 0, "")
				It("should return err", func() {
					Expect(err != nil).To(Equal(true))
				})
			})
		})

		Describe("with valid move", func() {
			Context("when facing north", func() {
				robot := placeRobot(0, 3, "NORTH")
				x, y, f, err := robot.Perform("MOVE", 0, 0, "")
				It("should be successful", func() {
					Expect(err == nil).To(Equal(true))
					Expect(x == 0).To(Equal(true))
					Expect(y == 4).To(Equal(true))
					Expect(f == "NORTH").To(Equal(true))
				})
			})

			Context("when facing south", func() {
				robot := placeRobot(0, 1, "SOUTH")
				_, _, _, err := robot.Perform("MOVE", 0, 0, "")
				It("should be successful", func() {
					Expect(err == nil).To(Equal(true))
					Expect(robot.X == 0).To(Equal(true))
					Expect(robot.Y == 0).To(Equal(true))
					Expect(robot.F == "SOUTH").To(Equal(true))
				})
			})

			Context("when facing east", func() {
				robot := placeRobot(3, 0, "EAST")
				x, y, f, err := robot.Perform("MOVE", 0, 0, "")
				It("should be successful", func() {
					Expect(err == nil).To(Equal(true))
					Expect(x == 4).To(Equal(true))
					Expect(y == 0).To(Equal(true))
					Expect(f == "EAST").To(Equal(true))
				})
			})

			Context("when facing west", func() {
				robot := placeRobot(1, 0, "WEST")
				_, _, _, err := robot.Perform("MOVE", 0, 0, "")
				It("should be successful", func() {
					Expect(err == nil).To(Equal(true))
					Expect(robot.X == 0).To(Equal(true))
					Expect(robot.Y == 0).To(Equal(true))
					Expect(robot.F == "WEST").To(Equal(true))
				})
			})
		})
	})

	Describe("Left", func() {
		Context("when robot is not place on the table", func() {
			robot := initialiseRobot()
			_, _, _, err := robot.Perform("LEFT", 0, 0, "")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Context("when robot facing North", func() {
			robot := placeRobot(0, 0, "NORTH")
			_, _, _, err := robot.Perform("LEFT", 0, 0, "")
			It("should face West", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "WEST").To(Equal(true))
			})
		})

		Context("when robot facing South", func() {
			robot := placeRobot(0, 0, "SOUTH")
			_, _, _, err := robot.Perform("LEFT", 0, 0, "")
			It("should face East", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "EAST").To(Equal(true))
			})
		})

		Context("when robot facing West", func() {
			robot := placeRobot(0, 0, "WEST")
			_, _, _, err := robot.Perform("LEFT", 0, 0, "")
			It("should face South", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "SOUTH").To(Equal(true))
			})
		})

		Context("when robot facing East", func() {
			robot := placeRobot(0, 0, "EAST")
			_, _, _, err := robot.Perform("LEFT", 0, 0, "")
			It("should face NORTH", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "NORTH").To(Equal(true))
			})
		})
	})

	Describe("Right Command", func() {
		Context("when robot is not place on the table", func() {
			robot := initialiseRobot()
			_, _, _, err := robot.Perform("RIGHT", 0, 0, "")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Context("when robot facing North", func() {
			robot := placeRobot(0, 0, "NORTH")
			_, _, _, err := robot.Perform("RIGHT", 0, 0, "")
			It("should face East", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "EAST").To(Equal(true))
			})
		})

		Context("when robot facing South", func() {
			robot := placeRobot(0, 0, "SOUTH")
			_, _, _, err := robot.Perform("RIGHT", 0, 0, "")
			It("should face West", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "WEST").To(Equal(true))
			})
		})

		Context("when robot facing West", func() {
			robot := placeRobot(0, 0, "WEST")
			_, _, _, err := robot.Perform("RIGHT", 0, 0, "")
			It("should face North", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "NORTH").To(Equal(true))
			})
		})

		Context("when robot facing East", func() {
			robot := placeRobot(0, 0, "EAST")
			_, _, _, err := robot.Perform("RIGHT", 0, 0, "")
			It("should face SOUTH", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "SOUTH").To(Equal(true))
			})
		})
	})

	Describe("Report Command", func() {
		Context("when robot is not place on the table", func() {
			robot := initialiseRobot()
			_, _, _, err := robot.Perform("REPORT", 0, 0, "")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Context("when robot is not place on the table", func() {
			robot := placeRobot(1, 1, "NORTH")
			x, y, f, err := robot.Perform("REPORT", 0, 0, "")
			It("should return err", func() {
				Expect(err == nil).To(Equal(true))
				Expect(x == 1).To(Equal(true))
				Expect(y == 1).To(Equal(true))
				Expect(f == "NORTH").To(Equal(true))
			})
		})
	})

	Describe("Initiate Robot", func() {
		Context("with invalid data", func() {
			_, err := NewRobot(-1, 0)
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})
		Context("with valid data", func() {
			robot, err := NewRobot(1, 3)
			It("should return err", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.Table.Width == 1).To(Equal(true))
				Expect(robot.Table.Length == 3).To(Equal(true))
			})
		})
	})
})

func placeRobot(x float32, y float32, f string) Robot {
	robot := initialiseRobot()
	robot.Place(x, y, f)
	return robot
}

func initialiseRobot() Robot {
	robot, _ := NewRobot(10, 10)
	return robot
}
