package playground

func validateBase(base string) bool {
	for _, ch := range base {
		if ch == '-' || ch == '+' {
			return false
		}
	}

	return true
}

func Pow(n, x int) int {
	if x == 0 {
		return 1
	}
	num := n
	for x > 1 {
		num *= n
		x--
	}

	return num
}

func GetNum(s string, ch rune) int {
	for i, v := range s {
		if ch == v {
			return i
		}
	}
	return -1
}

func AtoiBase(s string, base string) (num int) {
	if !validateBase(base) {
		return 0
	}

	b := len(base)
	p := len(s) - 1

	for _, ch := range s {
		n := GetNum(base, ch)

		num += n * Pow(b, p)

		p--
	}

	return num
}
