package main

type mFloat float64

func (m mFloat) divideWithFloat(divider int) mFloat {
	return (mFloat(divider) / m)
}
