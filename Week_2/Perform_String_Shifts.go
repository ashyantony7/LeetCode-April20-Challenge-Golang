package main

func stringShift(s string, shift [][]int) string {

	var totalShift, r int

	for i := 0; i < len(shift); i++ {
		totalShift += (shift[i][0]*2 - 1) * shift[i][1]
	}

	if totalShift > 0 {
		// left shift
		r = len(s) - (totalShift % len(s))
	} else {
		// right shift
		r = -(totalShift % len(s))
	}

	sRune := []rune(s)
	sRune = append(sRune[r:], sRune[:r]...)

	return string(sRune)

}
