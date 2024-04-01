package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
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
		cells:  make([][]bool, height),
	}
	for i := range maze.cells {
		maze.cells[i] = make([]bool, width)
	}
	return maze
}

func (m *Maze) Generate() {
	m.startPoint = [2]int{rand.Intn(m.width-2) + 1, rand.Intn(m.height-2) + 1}
	m.endPoint = [2]int{rand.Intn(m.width-2) + 1, rand.Intn(m.height-2) + 1}

	m.carvePassages(m.startPoint[0], m.startPoint[1])
}

func (m *Maze) carvePassages(currentX, currentY int) {
	directions := rand.Perm(4)
	for _, d := range directions {
		nextX, nextY := currentX, currentY
		switch d {
		case 0: // Up
			nextY -= 2
			if nextY < 0 || m.cells[nextY][nextX] {
				continue
			}
			m.cells[currentY-1][currentX] = true
		case 1: // Right
			nextX += 2
			if nextX >= m.width || m.cells[nextY][nextX] {
				continue
			}
			m.cells[currentY][currentX+1] = true
		case 2: // Down
			nextY += 2
			if nextY >= m.height || m.cells[nextY][nextX] {
				continue
			}
			m.cells[currentY+1][currentX] = true
		case 3: // Left
			nextX -= 2
			if nextX < 0 || m.cells[nextY][nextX] {
				continue
			}
			m.cells[currentY][currentX-1] = true
		}

		m.cells[nextY][nextX] = true
		m.PrintMaze()
		time.Sleep(1 * time.Second)
		m.carvePassages(nextX, nextY)
	}
}

func (m *Maze) PrintMaze() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println(m.String())
}

func (m *Maze) String() string {
	var output string
	for y := range m.cells {
		for x := range m.cells[y] {
			if x == 0 || x == m.width-1 || y == 0 || y == m.height-1 {
				output += "⬛"
			} else if x == m.startPoint[0] && y == m.startPoint[1] {
				output += "🟢"
			} else if x == m.endPoint[0] && y == m.endPoint[1] {
				output += "🔵"
			} else if m.cells[y][x] {
				output += "⬜️"
			} else {
				output += "️⬛"
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
	maze.PrintMaze()
}
