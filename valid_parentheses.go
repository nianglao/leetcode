package leetcode

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.
//
// Example 1:
//
// Input: "()"
// Output: true
// Example 2:
//
// Input: "()[]{}"
// Output: true
// Example 3:
//
// Input: "(]"
// Output: false
// Example 4:
//
// Input: "([)]"
// Output: false
// Example 5:
//
// Input: "{[]}"
// Output: true

import (
	"strings"
)

type Stack struct {
	Size int
	Data []string
}

func NewStack(i int) *Stack {
	var s Stack
	s.Size = 0
	s.Data = make([]string, i)
	return &s
}

func (s *Stack) Push(t string) {
	s.Data[s.Size] = t
	s.Size++
}

func (s *Stack) Pop() string {
	t := s.Data[s.Size-1]
	s.Size--
	return t
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	st := NewStack(len(s)/2 + 1)
	for _, c := range strings.Split(s, "") {
		switch c {
		case "{":
			st.Push("}")
		case "[":
			st.Push("]")
		case "(":
			st.Push(")")
		case "}", "]", ")":
			if st.Size == 0 {
				return false
			} else {
				if c != st.Pop() {
					return false
				}
			}
		}
	}
	return st.Size == 0
}
