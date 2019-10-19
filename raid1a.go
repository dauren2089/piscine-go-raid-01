package student

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {
	if x > 0 && y > 0 {
		if 0 < y && y < 3 {
			for j := 0; j < y; j++ {
				Example(x)
				z01.PrintRune(10)
			}
		} else {
			for j := 1; j <= y; j++ {
				if j == 1 || j == y {
					Example(x)
				} else {
					z01.PrintRune(10)
					CheckY(x)
				}
			}
		}
		z01.PrintRune(10)
	}
}

func Example(x int) {
	if 0 < x && x < 3 {
		for i := 0; i < x; i++ {
			z01.PrintRune(111)
		}
	} else {
		for i := 1; i <= x; i++ {
			if i == 1 || i == x {
				z01.PrintRune(111)
			} else {
				z01.PrintRune(45)
			}

		}
	}
}

func CheckY(x int) {

	for i := 1; i <= x; i++ {
		if i == 1 || i == x {
			z01.PrintRune(124)
		} else {
			z01.PrintRune(32)
		}
	}
}
