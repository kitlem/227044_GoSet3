package main

import (
	"fmt"
)

// Relationship Status
func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string]interface{}) string {
	// setting the variables as a boolean (to allow true and false logics)
	follows, _ := socialGraph[fromMember][toMember].(bool)
	followed_by, _ := socialGraph[toMember][fromMember].(bool)

	// conditions where if they both follow each other back
	if follows && followed_by {
		return fmt.Sprintf("%s and %s are friends", fromMember, toMember)
		// if fromMember follows toMember, but toMember doesnt follow back
	} else if follows {
		return fmt.Sprintf("%s follows %s", fromMember, toMember)
		// if fromMember follows toMember, but toMember does not follow fromMember
	} else if followed_by {
		return fmt.Sprintf("%s is followed by %s", fromMember, toMember)
	}

	// if neither follow each othber
	return fmt.Sprintf("%s and %s do not have a relationship", fromMember, toMember)
}

func main() {
	// Dictionaries where the true if person follows the person outside the dictionary , blank if not
	socialGraph := map[string]map[string]interface{}{
		"Aloy":   {"Erend": true, "Sylens": true},
		"Erend":  {"Aloy": true},
		"Sylens": {},
	}
	fmt.Println(relationshipStatus("Aloy", "Erend", socialGraph))
	fmt.Println(relationshipStatus("Aloy", "Sylens", socialGraph))
	fmt.Println(relationshipStatus("Sylens", "Aloy", socialGraph))
	fmt.Println(relationshipStatus("Erend", "Sylens", socialGraph))
}

// Tic Tac Toe

func ticTacToe(board [][]string) string {
	size := len(board)

	// Row and Column match
	// using var "i" to search the rows, 0 is used to begin
	for i := 0; i < size; i++ {
		if Equal(board[i]) {
			return board[i][0] // winner found when all matching elements are found in the slice (through rows only)
		}
		// Retrieving all elements in column
		column := []string{}
		// letting variable j search the columns
		for j := 0; j < size; j++ {
			column = append(column, board[j][i]) // winner found when all matching elements are found in the slice (vertical)
		}
		if Equal(column) {
			return column[0]
		}
	}

	// Diagonal match
	rightDiagonal := []string{}
	for i := 0; i < size; i++ {
		rightDiagonal = append(rightDiagonal, board[i][i]) // if a diagonal element has been found in the board
	}
	if Equal(rightDiagonal) {
		return rightDiagonal[0]
	}

	// Left-diagonal match (left to right), in order to check both sides
	leftDiagonal := []string{}
	for i := 0; i < size; i++ {
		leftDiagonal = append(leftDiagonal, board[i][size-i-1])
	}
	if Equal(leftDiagonal) {
		return leftDiagonal[0]
	}
	// if no matches are found (diagonal, vertical, lateral) matches, no winner is used
	return "No Winner"
}

// Checking if all elements are the same and not empty
func Equal(arr []string) bool {
	if len(arr) == 0 || arr[0] == "" {
		return false // if slice is empty, returns false
	}
	// If the elements are not the same, return false
	// ignores the index, and uses val to search within the arrray
	for _, val := range arr {
		if val != arr[0] {
			return false
		}
	}
	return true
}

func main() {
	// Manually draw the board (please input X and O with capital) using slices
	// each row is a slice and each element in a row is a string
	board := [][]string{
		{"X", "O", "O"},
		{"O", "X", "O"},
		{"O", "X", "O"},
	}
	// Winner printed on the board
	fmt.Println("The Winner is:", ticTacToe(board))
}

// ETA

func eta(firstStop string, secondStop string, routeMap map[string]map[string]int) int {
	Total_Time := 0           // total amount of time for the journey
	Current_Stop := firstStop // setting the default as first stop

	// this for loop is used when the next approach
	for Current_Stop != secondStop {
		// these variables are called in the next function
		Next_Stop, Travel_Time := nextLeg(Current_Stop, routeMap)
		// if next_stop is an empty variable returns -1 as an error message in main
		if Next_Stop == "" {
			return -1
		}
		// adding travel time to total time and moving to the next stop
		Total_Time += Travel_Time
		Current_Stop = Next_Stop
	}
	return Total_Time
}

func nextLeg(Current_Stop string, routeMap map[string]map[string]int) (string, int) {
	// iterating which are the next possible stops + time
	for next, time := range routeMap[Current_Stop] {
		return next, time
	}
	// if empty, returns 0
	return "", 0
}

func main() {
	// making route map with values for each destination (using nested map)
	routeMap := map[string]map[string]int{
		"A": {"B": 5},
		"B": {"C": 8},
		"C": {"D": 6},
		"D": {"A": 8},
	}

	// requires user input
	firstStop := "A"
	secondStop := "C"
	time := eta(firstStop, secondStop, routeMap)
	// if blank, prints error message
	if time == -1 {
		fmt.Println("Invalid route")
	} else {
		fmt.Printf("The time of the journey to go from %s to %s takes %d Minutes\n", firstStop, secondStop, time)
	}
}
