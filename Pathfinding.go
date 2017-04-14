package main

import (
	"fmt"
	"bufio"
	"os"
)

// Vertex is a struct used for every cell (character) in the maze
// It hold fields needed to perform Dijkstra's pathfinding algorithm
type Vertex struct {
	cellChar string	
	cost int
	lastVertexRow int
	lastVertexCol int
	isOpen bool
	row int
	col int
}

func main() {
	maze := readMaze()
	cheapestPath := calcShortestPath(maze[0:])
	printMaze(maze[0:])
	fmt.Println("The path cost:", cheapestPath, "HP")
}

// The purpose of calcShortestPath is to find the optimal path
// to the end of the maze. It also calculates the HP needed to get to
// the end of the maze, as well as draws in the path on the 
// maze using "*" characters
func calcShortestPath(maze [][]Vertex) int {
	var currVertex, destination, lastVertex *Vertex
	finished := false
	
	currVertex = findCheapestOpenVertex(maze[0:])	// Finds the starting point "S"
	
	for !finished {						// Keeps looping until the optimal path to the exit is found
		currVertex.isOpen = false		// Marks analyzed vertex as visited, so as to not visit it again
		if currVertex.cellChar == "G" {	// Set the finished flag, if the optimal route to exit is found
			destination = currVertex
			finished = true
		}
		
		updateTouchingVertices(currVertex, maze[0:])	// Calculates the costs of adjacent vertices
		
		currVertex = findCheapestOpenVertex(maze[0:])	// Analyzes the next cheapest vertex
	}
	
	lastVertex = &maze[destination.lastVertexRow][destination.lastVertexCol]	// Jumps to the last visited vertex from the exit vertex
	finished = false	// Resets the flag, so that it can be used in the upcoming loop
	for !finished {		// Keep looping	until the optimal path is drawn
		if lastVertex.cellChar != "S" && lastVertex.cellChar != "G" {
			lastVertex.cellChar = "*"
		}
		
		if lastVertex.lastVertexRow != -1 || lastVertex.lastVertexCol != -1 {		// If not the starting vertex
			lastVertex = &maze[lastVertex.lastVertexRow][lastVertex.lastVertexCol]	// Jump to the previous vertex
		} else {
			finished = true
		}
	}
	
	return destination.cost
}

// The purpose of updateTouchingVertices is to update the costs
// of the adjacent (from the current vertex) vertices
func updateTouchingVertices(currVertex *Vertex, maze [][]Vertex) {
	tempVertex := &maze[currVertex.row - 1][currVertex.col]	// Vertex up
	if tempVertex.cellChar != "#" && tempVertex.isOpen {	// If it's not a wall and it hasn't been visited
		updateTargetEstimate(currVertex, tempVertex)
	}
	
	tempVertex =  &maze[currVertex.row + 1][currVertex.col]	// Vertex down
	if tempVertex.cellChar != "#" && tempVertex.isOpen {
		updateTargetEstimate(currVertex, tempVertex)
	}
	
	tempVertex =  &maze[currVertex.row][currVertex.col - 1]	// Vertex left
	if tempVertex.cellChar != "#" && tempVertex.isOpen {
		updateTargetEstimate(currVertex, tempVertex)
	}
	
	tempVertex =  &maze[currVertex.row][currVertex.col + 1]	// Vertex right
	if tempVertex.cellChar != "#" && tempVertex.isOpen {
		updateTargetEstimate(currVertex, tempVertex)
	}
}

// The purpose of updateTargetEstimate is to update the cost
// of a cell by adding the previous costs to the current vertex
// and add the vertex char value ("m" = 11, " " = 1) to the sum
func updateTargetEstimate(currVertex *Vertex, targetVertex *Vertex) {
	var newCost int	// Will store the new total cost to this vertex
	if targetVertex.cellChar == "m" {
		newCost = currVertex.cost + 11
	} else {
		newCost = currVertex.cost + 1
	}
	
	// If it hasn't been checked yet or it's cost is higher than the cost of the new route, change its cost and the vertex from which to get there
	if targetVertex.cost == -1 || targetVertex.cost > newCost {
		targetVertex.cost = newCost
		targetVertex.lastVertexRow = currVertex.row
		targetVertex.lastVertexCol = currVertex.col
	}
}

// The purpose of findCheapestOpenVertex is to loop through the maze
// until the vertex (that hasn't been visited) with the cheapest path
// cost (so far) is found.
func findCheapestOpenVertex(maze [][]Vertex) *Vertex {
	var returnVertex Vertex
	var returnPointer *Vertex

	for row := 0; row < len(maze); row++ {
		for col := 0; col < len(maze[row]); col++ {
			// If it's not a wall ("#") and it hasn't been visited and it has been an adjacent vertex to one of the visited vertices
			if maze[row][col].cellChar != "#" && maze[row][col].isOpen && maze[row][col].cost >= 0 {
				if returnVertex == (Vertex{}) {		// The first vertex to meet the conditions is set as the return vertex
					returnPointer = &maze[row][col]	// Pointer to the vertex in the maze
					returnVertex = *returnPointer	// Contents in the vertex (its struct)
				} else if maze[row][col].cost < returnVertex.cost {	// If the costs are lower than the first vertex, then replace the return vertex
					returnPointer = &maze[row][col]	// Pointer to the vertex in the maze
					returnVertex = *returnPointer	// Contents in the vertex (its struct)
				}
			}
		}
	}

	return returnPointer
}

// The purpose of readMaze is to read the values from input.txt
// and construct a multidimensional array of Vertex, which it then returns
func readMaze() [][]Vertex {
	file, err := os.Open("input.txt")	// Try to open a file
    if err != nil {
        fmt.Println("Error: Can't read the file")
    }
    defer file.Close()	// Close the file only after the readMaze() function has executed
	
	row := 0
	maze := make([][]Vertex, 0)	// Create a multidimensional array of type Vertex
	
    scanner := bufio.NewScanner(file)	// Create a file scanner
    for scanner.Scan() {
		currentLine := scanner.Text()	// Read a single line
		
		tempRow := make([]Vertex, len(currentLine))	// Create a temporary array to hold our line
		for col, char := range currentLine {		// For each char in the line, create a Vertex
			currentChar := string(char)
			
			estimatedCost := -1
			if currentChar == "S" {
				estimatedCost = 0		// Starting point's cost is 0
			}
			
			// Create a Vertex struct and put it into the temporary row array
			tempRow[col] = Vertex {
				cellChar: 		currentChar,
				cost: 			estimatedCost,
				lastVertexRow: 	-1,
				lastVertexCol:	-1,
				isOpen: 		true,
				row:			row,
				col:			col,
			}
		}
		
		maze = append(maze, tempRow)	// Add the temp row onto the maze array	
		row++
    }

	return maze
}

// Purpose of printMaze is to print out the entire maze
func printMaze(maze [][]Vertex) {
	for row := 0; row < len(maze); row++ {
		for col := 0; col < len(maze[row]); col++ {
			fmt.Print(maze[row][col].cellChar)
		}
		fmt.Println()
	}
}