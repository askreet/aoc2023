package lib

func Combinations[T any](s []T, n int) [][]T {
	var cs [][]T

	for _, indexes := range combinationsIdx(len(s), n, 0, len(s)-1) {
		var c []T
		for _, i := range indexes {
			c = append(c, s[i])
		}
		cs = append(cs, c)
	}

	return cs
}

func combinationsIdx(total int, n int, start int, end int) [][]int {
	var cs [][]int

	for i := start; i <= end; i++ {
		head := []int{i}
		if n == 1 {
			cs = append(cs, head)
		} else {
			tails := combinationsIdx(total-i, n-1, i+1, end)
			for _, combination := range tails {
				cs = append(cs, append(head, combination...))
			}
		}
	}

	return cs
}
