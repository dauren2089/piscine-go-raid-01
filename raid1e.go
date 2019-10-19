package student

import (
	"github.com/01-edu/z01"
)

func Raid1e(x, y int) {
	if x > 0 && y > 0 {
		if 0 < y && y < 3 {
			for j := 0; j < y; j++ {
				FifthExample(x, y)
				z01.PrintRune(10)
			}
		} else {
			for j := 1; j <= y; j++ {
				if j == 1 || j == y {
					FifthCheckLastOrFirstY(x, j)
					z01.PrintRune(10)
				} else {
					FifthCheckY(x)
					z01.PrintRune(10)
				}
			}
		}
	}
}

func FifthExample(x, y int) {
	if 0 < x && x < 3 {
		for i := 0; i < x; i++ {
			z01.PrintRune(65)
		}
	} else {

		for i := 1; i <= x; i++ {
			if i == 1 && i == y { ///
				z01.PrintRune(65) // "/""
			} else if i == x && i == 1 {
				z01.PrintRune(65) // "\"
			} else if i == 1 { //3

				z01.PrintRune(47)
			} else if i == x {
				z01.PrintRune(67)
			} else {
				z01.PrintRune(66)
			}
		}
	}

}

func FifthCheckY(x int) {

	for i := 1; i <= x; i++ {
		if i == 1 || i == x {
			z01.PrintRune(66)
		} else {
			z01.PrintRune(32)
		}
	}
}

func FifthCheckLastOrFirstY(x, y int) {

	for i := 1; i <= x; i++ {
		if i == 1 && y != 1 {
			z01.PrintRune(67)
		} else if i == x && y != 1 {
			z01.PrintRune(65)
		} else if i == 1 && y == 1 {
			z01.PrintRune(65)
		} else if i == x && y == 1 {
			z01.PrintRune(67)
		} else {
			z01.PrintRune(66)
		}
	}

}
