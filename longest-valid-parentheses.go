type Stack struct {
    items []interface{}
}
func longestValidParentheses(s string) int {
    stack := Stack{}
    value:=0
    stack.Push(-1)
    for i:=0;i<len(s);i++ {
        if s[i]=='(' {
            stack.Push(i)
        } else {
            if !stack.isEmpty() {
                stack.Pop()
            }
            if !stack.isEmpty() {
                value=max(value,i-stack.peek().(int))
            } else {
                stack.Push(i)
            }
        }
    }
    return value
}
func (s *Stack) Push(item interface{}) {
    s.items = append(s.items,item)
}
func (s *Stack) Pop() interface{} {
    if len(s.items)==0 {
        return nil
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}
func (s *Stack) isEmpty() bool {
    return len(s.items)==0
}
func (s *Stack) peek() interface{} {
    if len(s.items)==0{
        return nil
    }
    item := s.items[len(s.items)-1]
    return item
}
func max(a,b int) int {
    if a>b {
        return a
    }
    return b
}