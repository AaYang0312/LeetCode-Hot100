package main

func isValid(s string) bool {
	var stc stack
	n := len(s)
	if n%2 != 0 {
		return false
	}
	for _, c := range s {
		if string(c) == "(" || string(c) == "[" || string(c) == "{" {
			stc.push(c)
			continue
		}
		if string(c) == ")" {
			if len(stc.queue) == 0 || string(stc.pop()) != "(" {
				return false
			}
			continue
		}
		if string(c) == "]" {
			if len(stc.queue) == 0 || string(stc.pop()) != "[" {
				return false
			}
			continue
		}
		if string(c) == "}" {
			if len(stc.queue) == 0 || string(stc.pop()) != "{" {
				return false
			}
			continue
		}
	}
	if len(stc.queue) != 0 {
		return false
	}
	return true
}

type stack struct {
	queue []int32
}

func (s *stack) push(c int32) {
	s.queue = append(s.queue, c)
	copy(s.queue, s.queue)
}
func (s *stack) pop() int32 {
	ans := s.queue[len(s.queue)-1]
	s.queue = s.queue[:len(s.queue)-1]
	copy(s.queue, s.queue)
	return ans
}
