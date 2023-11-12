func twoSum(nums []int, target int) []int {
    mop := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        complement := target - nums[i]
        if index, found := mop[complement]; found {
            return []int{index, i}
        }
        mop[nums[i]] = i
    }
    return []int{}

}