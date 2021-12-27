package main

import "fmt"

var grid = [][]byte{
	[]byte("########"),
	[]byte("#......#"),
	[]byte("#.###..#"),
	[]byte("#...#.##"),
	[]byte("#.#....#"),
	[]byte("########"),
}

type Coord struct {
	X, Y int
}
type StackCoord struct {
	PrevLoc []Coord
}

func (s *StackCoord) Push(c Coord) {
	s.PrevLoc = append(s.PrevLoc, c)
}
func (s *StackCoord) Pop() Coord {
	lastItem := s.PrevLoc[len(s.PrevLoc)-1]
	s.PrevLoc = s.PrevLoc[:len(s.PrevLoc)-1]
	return lastItem
}

var MOVE_UP = 0b001
var MOVE_RIGHT = 0b010
var MOVE_DOWN = 0b100

func CheckSurounding(curr Coord) int {
	//check up
	result := 0
	if grid[curr.Y-1][curr.X] == byte('.') {
		result = result ^ MOVE_UP
	}
	if grid[curr.Y][curr.X+1] == byte('.') {
		// kalo dah turn return 0
		if grid[curr.Y-1][curr.X] == byte('$') {
			return 0
		}
		result = result ^ MOVE_RIGHT
	}
	if grid[curr.Y+1][curr.X] == byte('.') {
		result = result ^ MOVE_DOWN
	}
	return result
}

var GlobalStack StackCoord

func PrintGrid(grid [][]byte) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}

func main() {
	GlobalStack = StackCoord{}
	CurrentCoord := Coord{1, 4}
	TreasurePossible := []Coord{}
	grid[CurrentCoord.Y][CurrentCoord.X] = byte('$')
	TreasurePossible = append(TreasurePossible, CurrentCoord)
	//grid2 := grid
	nextMove := CheckSurounding(CurrentCoord)
	counter := 0

	for nextMove != 0 || len(GlobalStack.PrevLoc) > 0 {
		counter++
		switch nextMove {
		case MOVE_UP:
			CurrentCoord.Y -= 1
		case MOVE_UP ^ MOVE_RIGHT:
			// put the right cell into stack, then move up
			leftCoord := Coord{CurrentCoord.X + 1, CurrentCoord.Y}
			GlobalStack.Push(leftCoord)
			// fmt.Println("Push Stack", leftCoord)
			CurrentCoord.Y -= 1
		case MOVE_RIGHT:
			CurrentCoord.X += 1
		case MOVE_DOWN ^ MOVE_RIGHT:
			// put down grid on stack, then move right
			leftCoord := Coord{CurrentCoord.X, CurrentCoord.Y + 1}
			GlobalStack.Push(leftCoord)
			// fmt.Println("Push Stack", leftCoord)
			CurrentCoord.X += 1
		case MOVE_DOWN:
			CurrentCoord.Y += 1
		case 0:
			CurrentCoord = GlobalStack.Pop()
			// fmt.Println("POP", CurrentCoord)
			// grid[CurrentCoord.Y][CurrentCoord.X] = byte('$')
		}
		grid[CurrentCoord.Y][CurrentCoord.X] = byte('$')
		TreasurePossible = append(TreasurePossible, CurrentCoord)
		nextMove = CheckSurounding(CurrentCoord)
	}
	PrintGrid(grid)
	for _, val := range TreasurePossible {
		fmt.Println(val.X, val.Y)
	}
	//grid[CurrentCoord.Y][CurrentCoord.X] = byte('X')
	//PrintGrid()
}
