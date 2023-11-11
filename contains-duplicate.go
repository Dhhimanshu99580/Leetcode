type hashSet struct {
	elements map[int]struct{}
  }
  
  func containsDuplicate(nums []int) bool {
	set := newHashSet()
	for i := 0; i < len(nums); i++ {
	  if set.Contains(nums[i]) {
		return true
	  } else {
		set.Add(nums[i])
	  }
	}
	return false
  }
  
  func newHashSet() *hashSet {
	return &hashSet{
	  elements: make(map[int]struct{}),
	}
  }
  
  func (s *hashSet) Add(element int) {
	s.elements[element] = struct{}{}
  }
  
  func (s *hashSet) Contains(element int) bool {
	_, ok := s.elements[element]
	return ok
  }