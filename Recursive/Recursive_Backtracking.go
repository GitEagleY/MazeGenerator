package main

import (
	"fmt"
	"math/rand"
)

type Maze struct {
	width      int
	height     int
	cells      [][]bool // 2D array representing cells of the maze
	startPoint [2]int   // coordinates
	endPoint   [2]int   // coordinates
}

func NewMaze(width, height int) *Maze {
	maze := &Maze{
		width:  width,
		height: height,
		cells:  make([][]bool, height), //height cells being initialized
	}
	for i := range maze.cells {
		maze.cells[i] = make([]bool, width) //width cells being initialized
	}
	return maze
}

func (m *Maze) Generate() {
	// Choose random coordinates for starting and ending points
	/*
		Subtracting 2 ensures that the generated random number is within the bounds of the maze
		but leaves a margin of 1 cell on each side.
	*/
	m.startPoint = [2]int{rand.Intn(m.width-2) + 1, rand.Intn(m.height-2) + 1}
	m.endPoint = [2]int{rand.Intn(m.width-2) + 1, rand.Intn(m.height-2) + 1}

	// Start carving passages from the starting point
	m.carvePassages(m.startPoint[0], m.startPoint[1])
}

// carvePassages carves passages in the maze using recursive backtracking algorithm
func (m *Maze) carvePassages(currentX, currentY int) {
	// Directions for movement: Up, Right, Down, Left
	directions := rand.Perm(4)
	for _, d := range directions {
		nextX, nextY := currentX, currentY
		switch d {
		case 0: // Up
			nextY -= 2
			// Check if the cell is within bounds and not already visited
			if nextY < 0 || m.cells[nextY][nextX] {
				continue
			}
			m.cells[currentY-1][currentX] = true // Mark the wall between current cell and next cell as passage
		case 1: // Right
			nextX += 2
			// Check if the cell is within bounds and not already visited
			if nextX >= m.width || m.cells[nextY][nextX] {
				continue
			}
			m.cells[currentY][currentX+1] = true // Mark the wall between current cell and next cell as passage
		case 2: // Down
			nextY += 2
			// Check if the cell is within bounds and not already visited
			if nextY >= m.height || m.cells[nextY][nextX] {
				continue
			}
			m.cells[currentY+1][currentX] = true // Mark the wall between current cell and next cell as passage
		case 3: // Left
			nextX -= 2
			// Check if the cell is within bounds and not already visited
			if nextX < 0 || m.cells[nextY][nextX] {
				continue
			}
			m.cells[currentY][currentX-1] = true // Mark the wall between current cell and next cell as passage
		}

		m.cells[nextY][nextX] = true  // Mark the next cell as visited
		m.carvePassages(nextX, nextY) // Recursively carve passages from the next cell
	}
}

// String returns the string representation of the maze
func (m *Maze) String() string {
	var output string
	// Loop through each cell in the maze
	for y := range m.cells {
		for x := range m.cells[y] {
			// Check for border cells, starting point, ending point, and passages
			if x == 0 || x == m.width-1 || y == 0 || y == m.height-1 {
				output += "⬛" // Border
			} else if x == m.startPoint[0] && y == m.startPoint[1] {
				output += "🟢" // Starting point
			} else if x == m.endPoint[0] && y == m.endPoint[1] {
				output += "🔵" // Ending point
			} else if m.cells[y][x] {
				output += "⬜️" // Passage
			} else {
				output += "⬛" // Wall
			}
		}
		output += "\n"
	}
	return output
}

func main() {
	var width, height int

	fmt.Println("Input width:")
	fmt.Scan(&width)
	fmt.Println("Input height:")
	fmt.Scan(&height)

	maze := NewMaze(width, height)
	maze.Generate()
	fmt.Println(maze.String())
}
