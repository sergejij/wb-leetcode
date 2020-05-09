package main

func markIsland(grid [][]byte, row int, col int) {
	grid[row][col] = '0'
	if row - 1 >= 0 && grid[row - 1][col] == '1' {
		markIsland(grid, row - 1, col)
	}
	if col - 1 >= 0 && grid[row][col - 1] == '1' {
		markIsland(grid, row, col - 1)
	}
	if row + 1 < len(grid) && grid[row + 1][col] == '1' {
		markIsland(grid, row + 1, col)
	}
	if col + 1 < len(grid[row]) && grid[row][col + 1] == '1' {
		markIsland(grid, row, col + 1)
	}
}

func numIslands(grid [][]byte) (res int) {
	numRow := len(grid)
	for i := 0; i < numRow; i++ {
		numCol := len(grid[i])
		for j := 0; j < numCol; j++ {
			if grid[i][j] == '1' {
				markIsland(grid, i, j)
				res++
			}
		}
	}
	return
}