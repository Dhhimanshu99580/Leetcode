func garbageCollection(garbage []string, travel []int) int {
    
	travel_time := 0
	   garbage_counter := make(map[rune][]int)
	   for i,item := range(garbage){
		   if i > 0{
			   travel_time+=travel[i-1] 
		   }
		   for _,letter := range(item){
			   if _,ok:= garbage_counter[letter]; !ok{
				   garbage_counter[letter] = []int{0,0}
			   }
			   garbage_counter[letter][0]++
			   garbage_counter[letter][1] = travel_time
		   }
	   }       
	   answer := 0
	   for _,value := range(garbage_counter){
		   counter := value[0]
		   time := value[1]
		   answer+=(counter+time)
	   }
	   return answer
   }