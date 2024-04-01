package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Maze struct {
	Width  int
	Height int
	Cells  [][]rune
}

const (
	Wall  = '⬛'
	Empty = '⬜'
)

func NewMaze(width, height int) *Maze {

	maze := &Maze{
		Width:  width,
		Height: height,
		Cells:  make([][]rune, height),
	}
	for i := range maze.Cells {
		maze.Cells[i] = make([]rune, width)
		for j := range maze.Cells[i] {
			maze.Cells[i][j] = Wall // All cells start as walls
		}
	}
	return maze
}

func (m *Maze) GenerateMaze() {
	// Choose a random starting point and mark it as part of the maze
	startX, startY := rand.Intn(m.Width), rand.Intn(m.Height)
	m.Cells[startY][startX] = Empty

	// Create a list to store frontier cells
	frontier := [][2]int{{startX, startY}}

	// Run Prim's Algorithm
	for len(frontier) > 0 {
		// Choose a random frontier cell
		idx := rand.Intn(len(frontier))
		current := frontier[idx]
		frontier = append(frontier[:idx], frontier[idx+1:]...)

		// Find neighbors of the current cell
		neighbors := m.getNeighbors(current[0], current[1])

		// Shuffle the neighbors to randomize the order
		rand.Shuffle(len(neighbors), func(i, j int) {
			neighbors[i], neighbors[j] = neighbors[j], neighbors[i]
		})

		// Connect the current cell to a random neighbor
		for _, neighbor := range neighbors {
			nx, ny := neighbor[0], neighbor[1]
			if m.isInBounds(nx, ny) && m.Cells[ny][nx] == Wall {
				// Open passage to the neighbor cell
				m.Cells[ny][nx] = Empty
				// Add the neighbor to the frontier
				frontier = append(frontier, neighbor)
				// Add one more cell in the same direction to make the passages wider
				m.Cells[ny+(current[1]-ny)/2][nx+(current[0]-nx)/2] = Empty

				time.Sleep(1 * time.Second)
				// Clear the screen
				fmt.Printf("\033[H\033[2J")
				// Print the maze
				fmt.Println(m.String())
			}
		}
	}
}

func (m *Maze) isInBounds(x, y int) bool {
	return x >= 0 && x < m.Width && y >= 0 && y < m.Height
}

// returns neighboring cells of a given cell
func (m *Maze) getNeighbors(x, y int) [][2]int {
	neighbors := make([][2]int, 0)
	//  offsets for neighboring cells
	directions := [][2]int{{-2, 0}, {2, 0}, {0, -2}, {0, 2}}
	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if m.isInBounds(nx, ny) {
			neighbors = append(neighbors, [2]int{nx, ny})
		}
	}
	return neighbors
}

func (m *Maze) String() string {
	var output strings.Builder

	for _, row := range m.Cells {
		for _, cell := range row {
			output.WriteRune(cell)
		}
		output.WriteRune('\n')
	}

	return output.String()
}

func main() {

	var width, height int

	fmt.Println("Input width:")
	fmt.Scan(&width)
	fmt.Println("Input height:")
	fmt.Scan(&height)

	maze := NewMaze(width, height)
	maze.GenerateMaze()
}
