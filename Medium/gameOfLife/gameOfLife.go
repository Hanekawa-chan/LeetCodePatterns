package gameOfLife

import "fmt"

func Example() {
	fmt.Println("\nGame Of Life")
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	fmt.Println("board =", board)
	gameOfLife(board)
	fmt.Println("result =", board)
}

func gameOfLife(board [][]int) {
	m, t := len(board), len(board[0])
	n := make([][]int, m)
	for i := range n {
		n[i] = make([]int, t)
		copy(n[i], board[i])
	}
	c := 0
	last := len(board) - 1
	last2 := len(board[0]) - 1
	if last > 0 && last2 > 0 {
		if board[0][0] == 0 {
			if board[0][1] == 1 && board[1][0] == 1 && board[1][1] == 1 {
				n[0][0] = 1
			}
		} else {
			if board[0][1] == 1 {
				c += 1
			}
			if board[1][1] == 1 {
				c += 1
			}
			if board[1][0] == 1 {
				c += 1
			}
			if c > 1 {
				n[0][0] = 1
			} else {
				n[0][0] = 0
			}
		}
		c = 0
		if board[0][last2] == 0 {
			if board[0][last2-1] == 1 && board[1][last2-1] == 1 && board[1][last2] == 1 {
				n[0][last2] = 1
			}
		} else {
			if board[0][last2-1] == 1 {
				c += 1
			}
			if board[1][last2] == 1 {
				c += 1
			}
			if board[1][last2-1] == 1 {
				c += 1
			}
			if c > 1 {
				n[0][last2] = 1
			} else {
				n[0][last2] = 0
			}
		}
		c = 0
		if board[last][last2] == 0 {
			if board[last][last2-1] == 1 && board[last-1][last2-1] == 1 && board[last-1][last2] == 1 {
				n[last][last2] = 1
			}
		} else {
			if board[last][last2-1] == 1 {
				c += 1
			}
			if board[last-1][last2] == 1 {
				c += 1
			}
			if board[last-1][last2-1] == 1 {
				c += 1
			}
			if c > 1 {
				n[last][last2] = 1
			} else {
				n[last][last2] = 0
			}
		}
		c = 0
		if board[last][0] == 0 {
			if board[last][1] == 1 && board[last-1][0] == 1 && board[last-1][1] == 1 {
				n[last][0] = 1
			}
		} else {
			if board[last][1] == 1 {
				c += 1
			}
			if board[last-1][1] == 1 {
				c += 1
			}
			if board[last-1][0] == 1 {
				c += 1
			}
			if c > 1 {
				n[last][0] = 1
			} else {
				n[last][0] = 0
			}
		}
		c = 0
		if last > 1 {
			for i := 1; i < last2; i++ {
				if board[0][i-1] == 1 {
					c += 1
				}
				if board[0][i+1] == 1 {
					c += 1
				}
				if board[1][i-1] == 1 {
					c += 1
				}
				if board[1][i] == 1 {
					c += 1
				}
				if board[1][i+1] == 1 {
					c += 1
				}
				if board[0][i] == 0 {
					if c == 3 {
						n[0][i] = 1
					}
				} else {
					if c == 3 || c == 2 {
						n[0][i] = 1
					} else {
						n[0][i] = 0
					}
				}
				c = 0
			}
			for i := 1; i < last; i++ {
				if board[i-1][0] == 1 {
					c += 1
				}
				if board[i+1][0] == 1 {
					c += 1
				}
				if board[i-1][1] == 1 {
					c += 1
				}
				if board[i][1] == 1 {
					c += 1
				}
				if board[i+1][1] == 1 {
					c += 1
				}
				if board[i][0] == 0 {
					if c == 3 {
						n[i][0] = 1
					}
				} else {
					if c == 3 || c == 2 {
						n[i][0] = 1
					} else {
						n[i][0] = 0
					}
				}
				c = 0
			}
			for i := 1; i < last; i++ {
				if board[i-1][last2] == 1 {
					c += 1
				}
				if board[i+1][last2] == 1 {
					c += 1
				}
				if board[i-1][last2-1] == 1 {
					c += 1
				}
				if board[i][last2-1] == 1 {
					c += 1
				}
				if board[i+1][last2-1] == 1 {
					c += 1
				}
				if board[i][last2] == 0 {
					if c == 3 {
						n[i][last2] = 1
					}
				} else {
					if c == 3 || c == 2 {
						n[i][last2] = 1
					} else {
						n[i][last2] = 0
					}
				}
				c = 0
			}
			for i := 1; i < last2; i++ {
				if board[last][i-1] == 1 {
					c += 1
				}
				if board[last][i+1] == 1 {
					c += 1
				}
				if board[last-1][i-1] == 1 {
					c += 1
				}
				if board[last-1][i] == 1 {
					c += 1
				}
				if board[last-1][i+1] == 1 {
					c += 1
				}
				if board[last][i] == 0 {
					if c == 3 {
						n[last][i] = 1
					}
				} else {
					if c == 3 || c == 2 {
						n[last][i] = 1
					} else {
						n[last][i] = 0
					}
				}
				c = 0
			}
			for i := 1; i < last; i++ {
				for j := 1; j < last2; j++ {
					if board[i-1][j-1] == 1 {
						c += 1
					}
					if board[i-1][j] == 1 {
						c += 1
					}
					if board[i-1][j+1] == 1 {
						c += 1
					}
					if board[i+1][j-1] == 1 {
						c += 1
					}
					if board[i+1][j] == 1 {
						c += 1
					}
					if board[i+1][j+1] == 1 {
						c += 1
					}
					if board[i][j-1] == 1 {
						c += 1
					}
					if board[i][j+1] == 1 {
						c += 1
					}
					if board[i][j] == 0 {
						if c == 3 {
							n[i][j] = 1
						}
					} else {
						if c == 3 || c == 2 {
							n[i][j] = 1
						} else {
							n[i][j] = 0
						}
					}
					c = 0
				}
			}
		}
		for i := range n {
			for j := range n[i] {
				board[i][j] = n[i][j]
			}
		}
	} else if last2 > 0 {
		n[0][0] = 0
		n[0][last2] = 0
		for i := 1; i < last2; i++ {
			if board[0][i-1] == 1 {
				c += 1
			}
			if board[0][i+1] == 1 {
				c += 1
			}
			if board[0][i] == 1 {
				if c == 2 {
					n[0][i] = 1
				} else {
					n[0][i] = 0
				}
			}
			c = 0
		}
		for i := range n {
			for j := range n[i] {
				board[i][j] = n[i][j]
			}
		}
	} else if last > 0 {
		n[0][0] = 0
		n[last][0] = 0
		for i := 1; i < last; i++ {
			if board[i-1][0] == 1 {
				c += 1
			}
			if board[i+1][0] == 1 {
				c += 1
			}
			if board[i][0] == 1 {
				if c == 2 {
					n[i][0] = 1
				} else {
					n[i][0] = 0
				}
			}
			c = 0
		}
		for i := range n {
			for j := range n[i] {
				board[i][j] = n[i][j]
			}
		}
	} else {
		board[0][0] = 0
	}
}
