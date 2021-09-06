package string

func isValid(s string) bool {
	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}

	var check string
	for _, value := range s {
		switch value {
		case '[', '(', '{':
			check += string(value)
		case ']':
			if len(check) > 0 && check[len(check)-1] == '[' {
				check = check[:len(check)-1]
			} else {
				return false
			}
		case ')':
			if len(check) > 0 && check[len(check)-1] == '(' {
				check = check[:len(check)-1]
			} else {
				return false
			}
		case '}':
			if len(check) > 0 && check[len(check)-1] == '{' {
				check = check[:len(check)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}

	return len(check) == 0
}
