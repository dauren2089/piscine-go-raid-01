package student

import (
	"github.com/01-edu/z01"
)

var Print = z01.PrintRune

var LetterA = rune(65)
var LetterB = rune(66)
var LetterC = rune(67)
var Space = rune(32)
var NewLine = rune(10)

func Raid1f(x, y int) {

	if IsEmpty(x, y) {
		if IsFirstOrSecond(y) {
			for j := 0; j < y; j++ {
				FirstOrSecondYRow(x, y)
				Print(NewLine)
			}
		} else {
			for j := 1; j <= y; j++ {
				if IsFirstOrLastRow(j, y) {
					FirstOrLastRow(x, j)
				} else {
					MiddleRow(x)
				}
				Print(NewLine)
			}
		}
	}
}

func FirstOrSecondYRow(x, y int) {

	if IsFirstOrSecond(x) {
		for i := 0; i < x; i++ {
			Print(LetterA)
		}
	} else {
		for i := 1; i <= x; i++ {
			if IsFirstRowAndFirstElement(i, y) {
				Print(LetterA)
			} else if IsFirstRowAndLastElement(i, x) {
				Print(LetterA)
			} else {
				Print(LetterB)
			}
		}
	}

}

func MiddleRow(x int) {

	for i := 1; i <= x; i++ {
		if IsFirstOrLastRow(i, x) {
			Print(LetterB)
		} else {
			Print(Space)
		}
	}
}

func FirstOrLastRow(x, y int) {
	for i := 1; i <= x; i++ {
		if i == 1 && y != 1 {
			Print(LetterC)
		} else if i == x && y != 1 {
			Print(LetterA)
		} else if i == 1 && y == 1 {
			Print(LetterA)
		} else if i == x && y == 1 {
			Print(LetterC)
		} else {
			Print(LetterB)
		}
	}

}

func IsEmpty(x, y int) bool {
	if x > 0 && y > 0 {
		return true
	} else {
		return false
	}
}

func IsFirstOrSecond(value int) bool {
	if 0 < value && value < 3 {
		return true
	} else {
		return false
	}
}

func IsFirstOrLastRow(j, y int) bool {
	if j == 1 || j == y {
		return true
	} else {
		return false
	}
}

func IsFirstRowAndFirstElement(i, y int) bool {
	if i == 1 && i == y {
		return true
	} else {
		return false
	}
}

func IsFirstRowAndLastElement(i, x int) bool {
	if i == x && i == 1 {
		return true
	} else {
		return false
	}
}
