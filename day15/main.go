package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Pair represents a coordinate in the warehouse.
type Pair struct {
	r int
	c int
}

// BigBox represents a box in the BigWarehouse with two parts: left and right.
type BigBox struct {
	left  Pair
	right Pair
}

// BigWarehouse represents the larger warehouse with robot, walls, and boxes.
type BigWarehouse struct {
	moveSeq  string
	boxes    map[BigBox]struct{}
	boxParts map[Pair]BigBox
	robot    Pair
	walls    map[Pair]struct{}
	width    int
	height   int
}

// Print prints the state of the BigWarehouse to the console.
func (b BigWarehouse) Print() {
	for r := 0; r < b.height; r++ {
		for c := 0; c < b.width; c++ {
			p := Pair{r, c}
			if p.r == b.robot.r && p.c == b.robot.c {
				fmt.Print("@")
				continue
			}
			if _, ok := b.walls[p]; ok {
				fmt.Print("#")
				continue
			}
			if bb, ok := b.boxParts[p]; ok {
				if p == bb.left {
					fmt.Print("[")
				} else {
					fmt.Print("]")
				}
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

// NewBigHouse transforms a regular Warehouse into a BigWarehouse.
func NewBigHouse(house Warehouse) BigWarehouse {
	var newHouse BigWarehouse
	newHouse.height = house.height
	newHouse.width = house.width * 2
	newHouse.moveSeq = house.moveSeq
	newHouse.boxParts = make(map[Pair]BigBox)
	newHouse.walls = make(map[Pair]struct{})
	newHouse.boxes = make(map[BigBox]struct{})

	// Transform robot position and walls for the BigWarehouse.
	newHouse.robot = Pair{house.robot.r, house.robot.c * 2}
	for wall := range house.walls {
		wall1 := Pair{wall.r, wall.c * 2}
		wall2 := Pair{wall1.r, wall1.c + 1}
		newHouse.walls[wall1] = struct{}{}
		newHouse.walls[wall2] = struct{}{}
	}

	// Transform boxes into BigBoxes.
	for box := range house.boxes {
		left := Pair{box.r, box.c * 2}
		right := Pair{box.r, left.c + 1}
		bigBox := BigBox{left, right}
		newHouse.boxParts[left] = bigBox
		newHouse.boxParts[right] = bigBox
		newHouse.boxes[bigBox] = struct{}{}
	}

	return newHouse
}

// Warehouse represents the original warehouse layout.
type Warehouse struct {
	moveSeq string
	boxes   map[Pair]struct{}
	robot   Pair
	walls   map[Pair]struct{}
	width   int
	height  int
}

// ToString converts the Warehouse state into a string for easier visualization.
func (w Warehouse) ToString() string {
	var ret string
	for r := 0; r < w.height; r++ {
		for c := 0; c < w.width; c++ {
			p := Pair{r, c}
			if p.r == w.robot.r && p.c == w.robot.c {
				ret += "@"
				continue
			}
			if _, ok := w.walls[p]; ok {
				ret += "#"
				continue
			}
			if _, ok := w.boxes[p]; ok {
				ret += "O"
				continue
			}
			ret += "."
		}
		ret += "\n"
	}
	return ret
}

// Print prints the Warehouse state to the console.
func (w Warehouse) Print() {
	for r := 0; r < w.height; r++ {
		for c := 0; c < w.width; c++ {
			p := Pair{r, c}
			if p.r == w.robot.r && p.c == w.robot.c {
				fmt.Print("@")
				continue
			}
			if _, ok := w.walls[p]; ok {
				fmt.Print("#")
				continue
			}
			if _, ok := w.boxes[p]; ok {
				fmt.Print("O")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

// readData reads the warehouse data from a file and returns a Warehouse object.
func readData(fileName string) Warehouse {
	scanner := bufio.NewScanner(strings.NewReader(fileName))
	var house Warehouse
	house.boxes = make(map[Pair]struct{})
	house.walls = make(map[Pair]struct{})
	scanner.Scan()

	// Handle top wall.
	topWalls := scanner.Text()
	for i := 0; i < len(topWalls); i++ {
		house.walls[Pair{0, i}] = struct{}{}
	}

	// Handle the warehouse layout.
	r := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line == topWalls {
			for i := 0; i < len(topWalls); i++ {
				house.walls[Pair{r, i}] = struct{}{}
			}
			break
		}
		chars := []rune(line)
		for col, char := range chars {
			switch char {
			case '#':
				house.walls[Pair{r, col}] = struct{}{}
			case 'O':
				house.boxes[Pair{r, col}] = struct{}{}
			case '@':
				house.robot = Pair{r, col}
			}
		}
		r++
	}

	// Set warehouse dimensions.
	house.height = r + 1
	house.width = len(topWalls)

	// Read the movement sequence.
	scanner.Scan()
	for scanner.Scan() {
		house.moveSeq += scanner.Text()
	}

	return house
}

// getNextPair calculates the next pair of coordinates based on the direction.
func getNextPair(p Pair, dir rune) Pair {
	nextPair := Pair{p.r, p.c}
	switch dir {
	case '^':
		nextPair.r--
	case '>':
		nextPair.c++
	case 'v':
		nextPair.r++
	case '<':
		nextPair.c--
	}
	return nextPair
}

// canBoxMove checks if a box can move in the given direction.
func canBoxMove(house *Warehouse, box Pair, dir rune) bool {
	nextCoords := getNextPair(box, dir)
	if _, ok := house.walls[nextCoords]; ok {
		return false // Wall encountered.
	}
	if _, ok := house.boxes[nextCoords]; ok {
		return canBoxMove(house, nextCoords, dir) // Check if the box can move further.
	}
	return true // No obstruction.
}

// moveBoxes moves a box to a new location.
func moveBoxes(house *Warehouse, box Pair, dir rune) {
	nextCoords := getNextPair(box, dir)
	if _, ok := house.boxes[nextCoords]; ok {
		moveBoxes(house, nextCoords, dir)
	}
	delete(house.boxes, box)
	house.boxes[nextCoords] = struct{}{}
}

// move moves the robot or a box in the specified direction.
func move(house *Warehouse, dir rune) {
	nextCoords := getNextPair(house.robot, dir)
	if _, ok := house.walls[nextCoords]; ok {
		return // Wall in front.
	}
	if _, ok := house.boxes[nextCoords]; ok {
		if canBoxMove(house, nextCoords, dir) { // Move the box if possible.
			moveBoxes(house, nextCoords, dir)
		} else {
			return // Cannot move box.
		}
	}
	house.robot = nextCoords // Move the robot.
}

// solveP1 solves part 1 of the problem.
func solveP1(house Warehouse) int {
	for _, char := range house.moveSeq {
		move(&house, char)
	}

	ret := 0
	for key := range house.boxes {
		ret += 100*key.r + key.c
	}

	return ret
}

// canBigBoxMove checks if a BigBox can move in the specified direction.
func canBigBoxMove(house *BigWarehouse, side Pair, dir rune) bool {
	canMove := true
	bb := house.boxParts[side]
	left, right := bb.left, bb.right
	leftNext := getNextPair(left, dir)
	rightNext := getNextPair(right, dir)

	// Check if walls block movement.
	if _, lOK := house.walls[leftNext]; lOK {
		canMove = false
	}
	if _, rOK := house.walls[rightNext]; rOK {
		canMove = false
	}

	// Check if the BigBox can move in the specified direction.
	if dir == '<' {
		if _, lOK := house.boxParts[leftNext]; lOK {
			canMove = canBigBoxMove(house, leftNext, dir)
		}
	} else if dir == '>' {
		if _, rOK := house.boxParts[rightNext]; rOK {
			canMove = canBigBoxMove(house, rightNext, dir)
		}
	} else {
		// Up or down movement.
		bbL, lOk := house.boxParts[leftNext]
		bbR, rOk := house.boxParts[rightNext]

		if lOk {
			canMove = canMove && canBigBoxMove(house, leftNext, dir)
		}
		if rOk && bbL != bbR {
			canMove = canMove && canBigBoxMove(house, rightNext, dir)
		}
	}
	return canMove
}

// bigBoxMove moves a BigBox in the specified direction.
func bigBoxMove(house *BigWarehouse, side Pair, dir rune) {
	bb := house.boxParts[side]
	left, right := bb.left, bb.right
	leftNext, rightNext := getNextPair(left, dir), getNextPair(right, dir)

	if dir == '<' {
		if _, lOK := house.boxParts[leftNext]; lOK {
			bigBoxMove(house, leftNext, dir)
		}
		delete(house.boxes, bb)
		delete(house.boxParts, left)
		delete(house.boxParts, right)

		bb.right = bb.left
		bb.left = leftNext
		house.boxes[bb] = struct{}{}
		house.boxParts[left] = bb
		house.boxParts[leftNext] = bb
		return
	}

	if dir == '>' {
		if _, rOK := house.boxParts[rightNext]; rOK {
			bigBoxMove(house, rightNext, dir)
		}
		delete(house.boxes, bb)
		delete(house.boxParts, left)
		delete(house.boxParts, right)

		bb.left = bb.right
		bb.right = rightNext

		house.boxes[bb] = struct{}{}
		house.boxParts[right] = bb
		house.boxParts[rightNext] = bb
		return
	}

	// For vertical movement (up or down).
	bbL, lOK := house.boxParts[leftNext]
	bbR, rOk := house.boxParts[rightNext]

	if lOK {
		bigBoxMove(house, leftNext, dir)
	}
	if rOk && bbL != bbR {
		bigBoxMove(house, rightNext, dir)
	}
	delete(house.boxes, bb)
	delete(house.boxParts, left)
	delete(house.boxParts, right)

	bb.left = leftNext
	bb.right = rightNext

	house.boxes[bb] = struct{}{}
	house.boxParts[leftNext] = bb
	house.boxParts[rightNext] = bb
}

// bigMove moves the robot or a BigBox in the specified direction.
func bigMove(house *BigWarehouse, dir rune) {
	nextCoords := getNextPair(house.robot, dir)

	if _, ok := house.walls[nextCoords]; ok {
		return // Wall in front.
	}

	if _, ok := house.boxParts[nextCoords]; ok {
		if canBigBoxMove(house, nextCoords, dir) {
			bigBoxMove(house, nextCoords, dir)
		} else {
			return // Cannot move BigBox.
		}
	}
	house.robot = nextCoords
}

// solveP2 solves part 2 of the problem with the BigWarehouse.
func solveP2(house BigWarehouse) int {
	for _, char := range house.moveSeq {
		bigMove(&house, char)
	}

	sum := 0
	for bb := range house.boxes {
		sum += 100*bb.left.r + bb.left.c
	}

	return sum
}

// SolutionForPart1 calculates the solution for part 1.
func SolutionForPart1(input string) (int, error) {
	house := readData(input)
	return solveP1(house), nil
}

// SolutionForPart2 calculates the solution for part 2.
func SolutionForPart2(input string) (int, error) {
	house := readData(input)
	bigHouse := NewBigHouse(house)
	return solveP2(bigHouse), nil
}
