package main

import (
	"fmt"
)

func main() {
	xoBoard := [3][3]string{}
	player := "x"
	for {
		fmt.Println("player :", player)
		var row int
		fmt.Println("please enter row number, 0, 1 or 2")
		fmt.Scanln(&row)
		if row > 2 {
			fmt.Println("invalid position, try another one")
			continue
		}

		var column int
		fmt.Println("please enter column number, 0, 1 or 2")
		fmt.Scanln(&column)
		if column > 2 {
			fmt.Println("invalid position, try another one")
			continue
		}

		if xoBoard[row][column] == "" {
			xoBoard[row][column] = player
		} else {
			fmt.Println("invalid position, try another one")
			continue
		}

		xoBoard[row][column] = player
		fmt.Println(xoBoard[0])
		fmt.Println(xoBoard[1])
		fmt.Println(xoBoard[2])

		for i := 0; i < 3; i++ {
			if (xoBoard[i][0] == player && xoBoard[i][1] == player && xoBoard[i][2] == player) || (xoBoard[0][i] == player && xoBoard[1][i] == player && xoBoard[2][i] == player) {
				fmt.Println("Winner")
				return
			}

		}

		if (xoBoard[0][0] == player && xoBoard[0][0] == xoBoard[1][1] && xoBoard[0][0] == xoBoard[2][2]) || (xoBoard[0][2] == player && xoBoard[0][2] == xoBoard[1][1] && xoBoard[0][2] == xoBoard[2][0]) {
			fmt.Println("Winner")
			return
		}

		if player == "x" {
			player = "o"
		} else {
			player = "x"
		}

	}
}
