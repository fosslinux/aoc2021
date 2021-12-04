package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

type cell struct {
	num   int
	state bool
}

type scoredBoard struct {
	board [][]cell
	score int
	id    int
}

func orderBoards(nums []int, boards [][][]cell) []scoredBoard {
	orderedBoards := make([]scoredBoard, 0)
	for _, num := range nums {
		for i, board := range boards {
			// Mark number
			for j, row := range board {
				for k, cell := range row {
					if cell.num == num {
						boards[i][j][k].state = true
					}
				}
			}
			// Won?
			won := false
			// Check rows
			for _, row := range board {
				thisWon := true
				for _, cell := range row {
					if !cell.state {
						thisWon = false
					}
				}
				won = won || thisWon
			}
			// Check columns
			for i := 0; i < 5; i++ {
				thisWon := true
				for j := 0; j < 5; j++ {
					if !board[j][i].state {
						thisWon = false
					}
				}
				won = won || thisWon
			}
			// Score if won
			if won {
				alreadyDone := false
				for _, scoredBoard := range orderedBoards {
					if scoredBoard.id == i {
						alreadyDone = true
					}
				}
				if !alreadyDone {
					sum := 0
					for _, row := range board {
						for _, cell := range row {
							if !cell.state {
								sum += cell.num
							}
						}
					}
					scored := scoredBoard{board: board, score: sum * num, id: i}
					orderedBoards = append(orderedBoards, scored)
				}
			}
		}
	}
	return orderedBoards
}

func main() {
	bin, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", err))
	}
	in := s.Split(string(bin), "\n")
	in = in[:len(in)-1]

	sNums := s.Split(in[0], ",")
	nums := make([]int, len(sNums))
	for i, str := range sNums {
		nums[i], _ = strconv.Atoi(str)
	}
	in = in[2:]

	// Yucky board parsing
	boards := make([][][]cell, 0)
	for i := 0; i < len(in)-4; i += 6 {
		board := make([][]cell, 5)
		for j := 0; j < 5; j++ {
			board[j] = make([]cell, 5)
			row := s.Split(in[i+j], " ")
			k := 0
			for _, num := range row {
				if num != "" {
					iNum, _ := strconv.Atoi(num)
					board[j][k] = cell{num: iNum, state: false}
					k++
				}
			}
		}
		boards = append(boards, board)
	}

	if os.Args[2] == "a" {
		fmt.Println(orderBoards(nums, boards)[0].score)
	} else if os.Args[2] == "b" {
		ordered := orderBoards(nums, boards)
		fmt.Println(ordered[len(ordered)-1].score)
	}
}
