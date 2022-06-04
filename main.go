package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 600
	screenHeight = 600
	cellSize     = 200
)

var (
	gameOver    = false
	draw        = false
	playersMove = true
	aiMove      = false
	x           rl.Texture2D
	o           rl.Texture2D
	player      int
)

type Cell struct {
	t        rl.Texture2D
	rect     rl.Rectangle
	marked   bool
	charType string
}

func checkWinner(cells []Cell) int {
	for _, cell := range cells {
		if cell.marked {
			if cell.charType == "x" {
				return 1
			} else if cell.charType == "o" {
				return 2
			}
		}
	}
	return 0
}

func minimax(cells []Cell, player int) int {
	winner := checkWinner(cells)
	if winner == 1 {
		return 10
	} else if winner == 2 {
		return -10
	} else if winner == 0 {
		return 0
	}

	var best int
	if player == 1 {
		best = -1000
	} else {
		best = 1000
	}

	for i := 0; i < 9; i++ {
		if cells[i].marked == false {
			cells[i].marked = true
			if player == 1 {
				cells[i].charType = "x"
				score := minimax(cells, 2)
				if score > best {
					best = score
				}
			} else {
				cells[i].charType = "o"
				score := minimax(cells, 1)
				if score < best {
					best = score
				}
			}
			cells[i].marked = false
			cells[i].charType = ""
		}
	}
	return best
}

func main() {

	rl.InitWindow(screenWidth, screenHeight, "Tic Tac Toe")
	rl.SetMouseScale(1.0, 1.0)

	x = rl.LoadTexture("assets/x.png")
	o = rl.LoadTexture("assets/o.png")

	var positions int = 0
	board := make([]Cell, 9)
	for i := 0; i < 9; i++ {
		board[i].rect.Width = cellSize
		board[i].rect.Height = cellSize
		board[i].rect.X = float32(i%3) * cellSize
		board[i].rect.Y = float32(i/3) * cellSize
	}

	var mouseButtonPressed bool = false

	for !rl.WindowShouldClose() {
		if !gameOver && !draw {

			var mousePos rl.Vector2 = rl.GetMousePosition()

			// if mouse is pressed, check if cell is empty or not
			if mouseButtonPressed {
				for i := 0; i < 9; i++ {
					if rl.CheckCollisionPointRec(mousePos, board[i].rect) && !board[i].marked && playersMove {
						board[i].t = x
						board[i].marked = true
						board[i].charType = "x"
						playersMove = false
						aiMove = true
						positions++
					}
					mouseButtonPressed = false
				}
			}
			// ai's turn
			if aiMove {
				var bestScore int = -1000
				var bestPos int = 0
				for i := 0; i < 9; i++ {
					if board[i].marked == false {
						board[i].marked = true
						board[i].charType = "o"
						score := minimax(board, 1)
						if score > bestScore {
							bestScore = score
							bestPos = i
						}
						board[i].marked = false
						board[i].charType = ""
					}
				}
				board[bestPos].t = o
				board[bestPos].marked = true
				board[bestPos].charType = "o"
				aiMove = false
				playersMove = true
				positions++
			}

			// ai movement logic
			// if aiMove {
			// 	var bestMove int
			// 	if positions == 0 {
			// 		bestMove = 0
			// 	} else {
			// 		bestMove = minimax(board, 2)
			// 	}
			// 	board[bestMove].t = o
			// 	board[bestMove].marked = true
			// 	board[bestMove].charType = "o"
			// 	aiMove = false
			// 	playersMove = true
			// 	positions++
			// }

			// check if game is over
			if positions == 9 {
				draw = true
			} else {
				winner := checkWinner(board)
				if winner == 1 {
					gameOver = true
					rl.DrawText("Player 1 wins!", screenWidth/2-50, screenHeight/2-50, 50, rl.Red)
				} else if winner == 2 {
					gameOver = true
					rl.DrawText("Player 2 wins!", screenWidth/2-50, screenHeight/2-50, 50, rl.Red)
				}
			}

			// some of the following code is from experimenting with the minimax algorithm!!!!!!!!

			// ai places its mark on its turn,
			// if aiMove {
			// 	var best int
			// 	if player == 1 {
			// 		best = 1000
			// 	} else {
			// 		best = -1000
			// 	}
			// 	for i := 0; i < 9; i++ {
			// 		if board[i].marked == false {
			// 			board[i].marked = true
			// 			board[i].charType = "o"
			// 			score := minimax(board, 1)
			// 			if score < best {
			// 				best = score
			// 			}
			// 			board[i].marked = false
			// 			board[i].charType = ""
			// 		}
			// 	}
			// 	for i := 0; i < 9; i++ {
			// 		if board[i].marked == false {
			// 			if rl.CheckCollisionPointRec(mousePos, board[i].rect) {
			// 				board[i].t = o
			// 				board[i].marked = true
			// 				board[i].charType = "o"
			// 				playersMove = true
			// 				aiMove = false
			// 				positions++
			// 			}
			// 		}
			// 	}
			// }
			// if aiMove {
			// 	var bestMove int
			// 	var bestScore int = -1000
			// 	for i := 0; i < len(board); i++ {
			// 		if !board[i].marked {
			// 			board[i].marked = true
			// 			board[i].charType = "o"
			// 			score := minimax(board, 1)
			// 			if score > bestScore {
			// 				bestScore = score
			// 				bestMove = i
			// 			}
			// 			board[i].marked = false
			// 			board[i].charType = ""
			// 		}
			// 	}
			// 	board[bestMove].t = o
			// 	board[bestMove].marked = true
			// 	board[bestMove].charType = "o"
			// 	playersMove = true
			// 	aiMove = false
			// 	positions++
			// }

			// ai move if player's move is over
			// if aiMove {
			// 	var bestMove int
			// 	if player == 1 {
			// 		bestMove = minimax(board, 2)
			// 	} else {
			// 		bestMove = minimax(board, 1)
			// 	}
			// 	for i := 0; i < len(board); i++ {
			// 		if board[i].marked == false && bestMove == i {
			// 			board[i].t = o
			// 			board[i].marked = true
			// 			board[i].charType = "o"
			// 			aiMove = false
			// 			playersMove = true
			// 			positions++
			// 		}
			// 	}
		}

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseRightButton) {
			mouseButtonPressed = true
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// draw grid on board
		rl.DrawLine(0, screenHeight/3, screenWidth, screenHeight/3, rl.White)
		rl.DrawLine(0, screenHeight*2/3, screenWidth, screenHeight*2/3, rl.White)
		rl.DrawLine(screenWidth/3, 0, screenWidth/3, screenHeight, rl.White)
		rl.DrawLine(screenWidth*2/3, 0, screenWidth*2/3, screenHeight, rl.White)

		if !gameOver && !draw {
			for i := 0; i < 9; i++ {
				rl.DrawRectangle(int32(board[i].rect.X), int32(board[i].rect.Y), int32(board[i].rect.Width-5), int32(board[i].rect.Height-5), rl.DarkBlue)

				rl.DrawTexture(board[i].t, int32(board[i].rect.X+board[i].rect.Width/2-float32(board[i].t.Width)/2), int32(board[i].rect.Y+board[i].rect.Height/2-float32(board[i].t.Height)/2), rl.White)

			}
		}

		if gameOver {
			rl.DrawText("Game Over!", screenWidth/2-50, screenHeight/2-50, 50, rl.Red)
		}

		if draw {
			rl.DrawText("Draw!", screenWidth/2-50, screenHeight/2-50, 50, rl.Red)
		}

		rl.EndDrawing()
	}
	rl.UnloadTexture(x)
	rl.UnloadTexture(o)
	rl.CloseWindow()
}
