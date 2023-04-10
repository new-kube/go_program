package main

func new_line_id() {
	a, x, y := 0, 1, 2

	a = x + y

	a = x +
		y

	// + 不能换行。
	a = x
	+y
}

func main() {
	new_line_id()
}
