package main

import "github.com/01-edu/z01"

func main() {
	//QuadA(5, 3)
	//QuadB(0, 6)
	//QuadC(5, 5)
	//QuadD(5, 3)
	//QuadE(5, 3)
}

func QuadA(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for c := 1; c <= y; c++ {
		for i := 1; i <= x; i++ {
			if c == 1 && i == 1 {
				z01.PrintRune('o')
			} else if c == 1 && i == x {
				z01.PrintRune('o')
			} else if c == y && i == 1 {
				z01.PrintRune('o')
			} else if c == y && i == x {
				z01.PrintRune('o')
			} else if c == 1 && i > 1 && i < x || c == y && i > 1 && i < x {
				z01.PrintRune('-')
			} else if c > 1 && i == 1 && c < y || c > 1 && i == x && c < y {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadB(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for c := 1; c <= y; c++ {
		for i := 1; i <= x; i++ {
			if c == 1 && i == 1 {
				z01.PrintRune('/')
			} else if c == 1 && i == x {
				z01.PrintRune('\\')
			} else if c == y && i == 1 {
				z01.PrintRune('\\')
			} else if c == y && i == x {
				z01.PrintRune('/')
			} else if c == 1 && i > 1 && i < x || c == y && i > 1 && i < x {
				z01.PrintRune('*')
			} else if c > 1 && i == 1 && c < y || c > 1 && i == x && c < y {
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}

}

func QuadC(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for c := 1; c <= y; c++ {
		for i := 1; i <= x; i++ {
			if c == 1 && i == 1 {
				z01.PrintRune('A')
			} else if c == 1 && i == x {
				z01.PrintRune('A')
			} else if c == y && i == 1 {
				z01.PrintRune('C')
			} else if c == y && i == x {
				z01.PrintRune('C')
			} else if c == 1 && i > 1 && i < x || c == y && i > 1 && i < x {
				z01.PrintRune('B')
			} else if c > 1 && i == 1 && c < y || c > 1 && i == x && c < y {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadD(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for c := 1; c <= y; c++ {
		for i := 1; i <= x; i++ {
			if c == 1 && i == 1 {
				z01.PrintRune('A')
			} else if c == 1 && i == x {
				z01.PrintRune('C')
			} else if c == y && i == 1 {
				z01.PrintRune('A')
			} else if c == y && i == x {
				z01.PrintRune('C')
			} else if c == 1 && i > 1 && i < x || c == y && i > 1 && i < x {
				z01.PrintRune('B')
			} else if c > 1 && i == 1 && c < y || c > 1 && i == x && c < y {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadE(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for c := 1; c <= y; c++ {
		for i := 1; i <= x; i++ {
			if c == 1 && i == 1 {
				z01.PrintRune('A')
			} else if c == 1 && i == x {
				z01.PrintRune('C')
			} else if c == y && i == 1 {
				z01.PrintRune('C')
			} else if c == y && i == x {
				z01.PrintRune('A')
			} else if c == 1 && i > 1 && i < x || c == y && i > 1 && i < x {
				z01.PrintRune('B')
			} else if c > 1 && i == 1 && c < y || c > 1 && i == x && c < y {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
