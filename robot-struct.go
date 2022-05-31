package main

/*
In this exercise you will modify the file robot.go to create a robot that can move around the coordinates of a 2-dimenensional plane in either the x or y direction.
*/

import (
	"fmt"
	"math"
)

/* 1)
Define a type "direction" based on the built-in type int
*/
type direction int

/* 2)
Create an enumeration of type direction with named constants north, south, east, and west
Use iota to set the sequence
*/
const (
	north direction = iota
	south
	east
	west
)

/* 3)
Define a type named "robot" based on struct
Add the following fields:
name (string)
x, y (int)
direction (direction)
*/
type robot struct {
	name string
	x, y int
	dir  direction
}

/* 4)
Create a method of robot named "initialize" that takes an argument of type string
Assign the argument to the name field of the robot
*/
func (r *robot) initialize(newName string) {
	r.name = newName
}

/*  5)
Create a method of robot named "move" that takes an argument of type direction
Using a switch statement:
increment x if the direction is north,
increment y if the direction is east,
decrement x if the direction is west,
decrement y if the direction is south,
*/
func (r *robot) move(newDirection direction) {
	switch newDirection {
	case 0:
		r.y += 1
	case 1:
		r.y -= 1
	case 2:
		r.x += 1
	case 3:
		r.x -= 1
	}
}

/* 6)
Create a method of robot named "calculateDistance" that returns a value of type float64
Using math.Sqrt, calculate the distance from the center using
the formula for the hypotenuse: sqrt(x * x + y * y)
Return the value
*/
func (r *robot) calculateDistance() float64 {
	return math.Sqrt(float64(r.x*r.x + r.y*r.y))
}

/* 7)
In main():
Create a robot
Set the name to "tobor"
Move north, east, north, north
Print the coordinates and the distance from the center
The output sould look like this:
Location: (1, 3)
Distance from center: 3.162278
*/
func main() {
	robo := robot{}
	robo.initialize("tobor")
	robo.move(north)
	robo.move(east)
	robo.move(north)
	robo.move(north)
	fmt.Printf("Location: (%d, %d)\n", robo.x, robo.y)
	fmt.Printf("Distance from center: %f\n", robo.calculateDistance())
}
