import "container/list"

func isValid(s string) bool {
	L := list.New()
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '{', '(', '[':
			L.PushBack(s[i])
		default:
			if L.Len() == 0 {
				return false
			}
			c := L.Back().Value.(byte)
			switch s[i] {
			case '}':
				if c != '{' {
					return false
				}
			case ')':
				if c != '(' {
					return false
				}
			case ']':
				if c != '[' {
					return false
				}
			}
			L.Remove(L.Back())
		}
	}
	if L.Len() > 0 {
		return false
	}
	return true
}
