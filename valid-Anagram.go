func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		 return false
	 }
	 chars1 := []rune(strings.ToLower(s))
	 chars2 := []rune(strings.ToLower(t))
	 sort.SliceStable(chars1, func(i, j int) bool {
		 return chars1[i] < chars1[j]
	 })
	 sort.SliceStable(chars2, func(i, j int) bool {
		 return chars2[i] < chars2[j]
	 })
	 for index, value := range chars1 {
		 if chars2[index] != value {
			 return false
		 }
	 }
	 return true
 }