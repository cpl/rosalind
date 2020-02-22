package rosalind

func Rabbits(generations, pairs int) int {
	a, b := 1, 1
	for generation := 1; generation < generations; generation++ {
		a, b = b, b+(a*pairs)
	}
	return a
}
