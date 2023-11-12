func maxArea(height []int) int {
	i,j,maximum,temp:=0,len(height)-1,0,0
	for i<j {
		if height[i]<=height[j] {
			temp=height[i]*(j-i)
			maximum=max(maximum,temp)
			i++
		} else {
			temp=height[j]*(j-i)
			maximum = max(maximum,temp)
			j--
		}
	} 
	return maximum
 }
 func max(a, b int) int {
	 if a > b {
		 return a
	 }
	 return b
 }