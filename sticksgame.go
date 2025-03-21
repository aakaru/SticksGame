package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sticks := 12
	turn := 0
	player1 := []int{}
	player2 := []int{}
	fmt.Println("There are 12 sticks in all. Pick one, two or three sticks. The player who picks the last stick loses and the other one is the winner.")
	reader := bufio.NewReader(os.Stdin)
	for sticks > 0 {
		turn++

		var FloatnumSticks float64

		currentPlayer := turn % 2
		if currentPlayer == 1 {
			fmt.Println("It's Player 1's turn!")
		} else {
			fmt.Println("It's Player 2's turn!")
		}
		playerPicks, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			continue
		}
		trimmed := strings.TrimSpace(playerPicks)
		FloatnumSticks, err = strconv.ParseFloat(trimmed, 64)
		if err != nil {
			fmt.Println("Please pick a stick!")
			continue
		}
		numSticks := int(FloatnumSticks)
		if numSticks == 1 || numSticks == 2 || numSticks == 3 {
			sticks -= int(numSticks)
			if currentPlayer == 1 {
				player1 = append(player1, int(numSticks))
				fmt.Printf("You Picked: %d\n", numSticks)
				fmt.Printf("Remaining sticks: %d\n", sticks)
			} else {
				player2 = append(player2, int(numSticks))
				fmt.Printf("You Picked: %d\n", numSticks)
				fmt.Printf("Remaining sticks: %d\n", sticks)

			}

		} else if numSticks > 3 {
			fmt.Printf("Don't cheat!\n")
			fmt.Printf("Remaining sticks: %d\n", sticks)
			turn--
		} else if numSticks <= 0 {
			fmt.Println("You need to pick at least one stick!\n")
			fmt.Printf("Remaining sticks: %d\n", sticks)
			turn--
		}

	}
	if sticks <= 0 {
		currentPlayer := turn % 2
		if currentPlayer == 1 {
			fmt.Println("Player 1 picked the last stick! Player 2 is the Winner!")
		} else {
			fmt.Println("Player 2 picked the last stick! Player 1 is the Winner!")
		}
	}
	for {
		fmt.Println("Do you want to view the Scoreboard of the game? Y or N: ")
		yOrn, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading your input")
			continue
		}

		yOrn = strings.TrimSpace(yOrn)

		if yOrn == "Y" {

			fmt.Println("Player 1 picks:", player1)
			fmt.Println("Player 2 picks:", player2)
			break

		} else if yOrn == "N" {
			fmt.Println("Thanks for playing!")
			break
		} else {
			fmt.Println("Please choose from Y or N.")
		}

	}
}
