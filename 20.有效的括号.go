/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

 type ELement interface{}

func isValid(s string) bool {
	stk := NewStack()
	m := make(map[string] string)
	m[")"] = "("
	m["}"] = "{"
	m["]"] = "["
	for i := 0; i < len(s); i++ {
		if c, ok := m[string(s[i])]; ok {
			if c != stk.Pop() {
				return false
			}
		} else {
			stk.Push(string(s[i]))
		}
	}
	return stk.IsEmpty()
}

type stack struct {
	element []ELement
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Push(e ELement) {
	s.element = append(s.element, e)
}

func (s *stack) Pop() ELement {
	size := s.Size()
	if size == 0 {
		return nil
	}
	lastElement := s.element[size - 1]
	s.element[size - 1] = nil
	s.element = s.element[:size - 1]
	return lastElement
}

func (s *stack) IsEmpty() bool {
	if len(s.element) == 0 {
		return true
	}
	return false
}

func (s *stack) Size() int {
	return len(s.element)
}
