func productExceptSelf(nums []int) []int {
	slice := make([]int, len(nums))
   mul,count,index := 1, 0, 0
   for i:=0;i<len(nums);i++ {
	   if nums[i] == 0 {
		   count++
		   index=i
	   }
   }
   if count==1 {
	   for i:=0;i<len(nums);i++ {
		   if i==index {
			   continue
		   }
		   mul*=nums[i]
		   slice[i]=0
	   }
	   slice[index]=mul
   } else if(count<1) {
	   for i:=0;i<len(nums);i++ {
		   mul*=nums[i]
	   }
	   for i:=0;i<len(nums);i++ {
		   slice[i] = mul/nums[i]
	   }
   } else {
	   for i:=0;i<len(nums);i++ {
		   slice[i]=0
	   }
   }
   return slice
}