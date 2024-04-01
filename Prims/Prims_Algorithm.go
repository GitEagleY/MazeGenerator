package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Maze struct {
	width  int
	height int
	cells  [][]rune
}

func NewMaze(width, height int) *Maze {
	maze := &Maze{
		width:  width,
		height: height,
		cells:  make([][]rune, height),
	}
	for i := range maze.cells {
		maze.cells[i] = make([]rune, width)
		for j := range maze.cells[i] {
			if i%2 == 1 && j%2 == 1 {
				maze.cells[i][j] = ' ' // Empty space for odd indices
			} else {
				maze.cells[i][j] = '#' // Wall for even indices
			}
		}
	}
	return maze
}

func (m *Maze) Generate() {
	// Initialize the maze with walls
	for i := range m.cells {
		for j := range m.cells[i] {
			m.cells[i][j] = '#'
		}
	}

	// Choose a random starting point and mark it as part of the maze
	startX, startY := rand.Intn(m.width), rand.Intn(m.height)
	m.cells[startY][startX] = ' '

	// Create a list to store frontier cells
	frontier := [][2]int{{startX, startY}}

	// Run Prim's Algorithm
	for len(frontier) > 0 {
		// Choose a random frontier cell
		randomIndex := rand.Intn(len(frontier))
		current := frontier[randomIndex]
		//remove the element at randomIndex from the frontier slice.
		frontier = append(frontier[:randomIndex], frontier[randomIndex+1:]...)

		// Find neighbors of the current cell
		neighbors := m.getNeighbors(current[0], current[1])

		// Shuffle the neighbors to randomize the order
		rand.Shuffle(len(neighbors), func(i, j int) {
			neighbors[i], neighbors[j] = neighbors[j], neighbors[i]
		})

		// Connect the current cell to a random neighbor
		for _, neighbor := range neighbors {

			// Extract the x and y coordinates of the neighbor cell from the neighbor array.
			nx, ny := neighbor[0], neighbor[1]
			// Check if the neighbor cell (nx, ny) is within the bounds of the maze and if it contains a wall ('#').
			if m.isInBounds(nx, ny) && m.cells[ny][nx] == '#' {
				// Open passage to the neighbor cell
				m.cells[ny][nx] = ' '
				// Add the neighbor to the frontier
				frontier = append(frontier, neighbor)
				// Add one more cell in the same direction to make the passages wider
				m.cells[ny+(current[1]-ny)/2][nx+(current[0]-nx)/2] = ' '
			}
		}
	}
}

// checks if a cell is within the bounds of the maze
func (m *Maze) isInBounds(x, y int) bool {
	return x >= 0 && x < m.width && y >= 0 && y < m.height
}

// returns neighboring cells of a given cell
func (m *Maze) getNeighbors(x, y int) [][2]int {
	neighbors := make([][2]int, 0)

	// offsets for neighboring cells
	directions := [][2]int{{-2, 0}, {2, 0}, {0, -2}, {0, 2}}

	for _, direction := range directions {
		nx, ny := x+direction[0], y+direction[1]

		// Check if the neighboring cell is within the bounds of the maze
		if m.isInBounds(nx, ny) {
			// add the coordinates to the list of neighbors
			neighbors = append(neighbors, [2]int{nx, ny})
		}
	}

	return neighbors
}

func (m *Maze) AddTreasureAndTraps(numTreasures, numTraps int) {
	m.addItems(numTreasures, 'T') // Add treasures
	m.addItems(numTraps, 'D')     // Add traps
}

func (m *Maze) addItems(num int, item rune) {
	for i := 0; i < num; i++ {
		x, y := rand.Intn(m.width), rand.Intn(m.height)
		if m.cells[y][x] == ' ' {
			m.cells[y][x] = item
		}
	}
}

func (m *Maze) AddStartAndEnd() {
	m.cells[1][1] = 'S'                  // Start marker
	m.cells[m.height-2][m.width-2] = 'E' // End marker
}

// returns the string representation of the maze
func (m *Maze) String() string {
	var output strings.Builder // Initialize a string builder to efficiently build the output

	// Top border
	output.WriteString("#")                              // top left border
	output.WriteString(strings.Repeat("#", m.width*2-1)) //  top border
	output.WriteString("#\n")                            //  top right border

	// Maze cells
	for _, row := range m.cells {
		output.WriteString("#") //  left border
		for _, cell := range row {
			output.WriteRune(cell)  //  cell
			output.WriteString(" ") //  space between cells
		}
		output.WriteString("#\n") //  right border
	}

	// Bottom border
	output.WriteString("#")                              // bottom left border
	output.WriteString(strings.Repeat("#", m.width*2-1)) //  bottom border
	output.WriteString("#\n")                            // bottom right border

	return output.String() // Return string
}

func main() {
	fmt.Println("S=start E=end T=treasure D=danger")

	var width, height int
	fmt.Println("Input width:")
	fmt.Scan(&width)
	fmt.Println("Input height:")
	fmt.Scan(&height)

	maze := NewMaze(width, height)
	maze.Generate()
	maze.AddTreasureAndTraps(5, 5)
	maze.AddStartAndEnd()
	fmt.Println(maze.String())
}
