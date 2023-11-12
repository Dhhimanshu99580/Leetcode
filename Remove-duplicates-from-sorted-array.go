func removeDuplicates(nums []int) int {
    if len(nums) == 1 {
    return 1
  }


  firstPtr := 0
  secPtr := 1
  curr  := 0

  for i := 0; i < len(nums); i++ {
    if secPtr == len(nums) {
      break
    }

    if nums[firstPtr] == nums[secPtr] {
      secPtr++
      continue
    }

    if nums[firstPtr] != nums[secPtr] {
      nums[curr + 1] = nums[secPtr]
      firstPtr = secPtr
      secPtr++
    }

    curr++
  }

  return curr + 1
    
}