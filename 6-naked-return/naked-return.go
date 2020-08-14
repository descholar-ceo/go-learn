package main

func namedReturn(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
