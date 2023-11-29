func largestSubmatrix(matrix [][]int) int {
    
	m, n := len(matrix), len(matrix[0])
	
		
		heights := make([][]int, m)
		for i := range heights {
			heights[i] = make([]int, n)
			for j := 0; j < n; j++ {
				if matrix[i][j] == 1 {
					if i > 0 {
						heights[i][j] = heights[i-1][j] + 1
					} else {
						heights[i][j] = 1
					}
				}
			}
		}
	
		
		for i := range heights {
			sort.Ints(heights[i])
		}
	
		
		maxArea := 0
		for i := m - 1; i >= 0; i-- {
			for j := 0; j < n; j++ {
				height := heights[i][j]
				width := n - j
				maxArea = max(maxArea, height*width)
			}
		}
	
		return maxArea
	}
	
	func max(a, b int) int {
		if a > b {
			return a
		}
		return b
	}