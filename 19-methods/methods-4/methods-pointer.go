package main

type mInt int

func (m mInt) absolute() mInt {
	if m < 0 {
		return (+m)
	}
	return m
}
