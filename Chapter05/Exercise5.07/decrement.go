package main

func decrement(i int) func() int {
	return func() int {
		i--
		return i
	}
}
