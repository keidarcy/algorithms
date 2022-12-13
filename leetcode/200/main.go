package main

import "fmt"

func main() {

	// init slice of byte slice golang
	// grid := [][]byte{
	// 	{'1', '1', '1', '1', '0'},
	// 	{'1', '1', '0', '1', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '0', '0', '0'},
	// }

	// [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	result := numIslands(grid)
	fmt.Printf("result: %d\n", result)
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				// dfs(grid, i, j)
				bfs(grid, i, j)
			}
		}
	}
	return count
}

// bfs
func bfs(grid [][]byte, r int, c int) {
	queue := [][]int{{r, c}}
	// queue = append(queue, []int{r, c})
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		r := point[0]
		c := point[1]
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != '1' {
			continue
		}
		grid[r][c] = '0'
		// down
		queue = append(queue, []int{r + 1, c})
		// left
		queue = append(queue, []int{r, c - 1})
		// up
		queue = append(queue, []int{r - 1, c})
		// right
		queue = append(queue, []int{r, c + 1})
	}
}

// dfs
func dfs(grid [][]byte, r int, c int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != '1' {
		return
	}

	grid[r][c] = '0'
	// down
	dfs(grid, r+1, c)
	// left
	dfs(grid, r, c-1)
	// up
	dfs(grid, r-1, c)
	// right
	dfs(grid, r, c+1)
}
