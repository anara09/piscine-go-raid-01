package main

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	// the fist case
	if y > 1 && x > 1 {
		// the first line
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune(10)
		//the second line
		for i := 0; i < y-2; i++ {
			z01.PrintRune('|')
			for i := 0; i < x-2; i++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('|')
			z01.PrintRune('\n')
		}
		//the third line
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune(10)
	}
	// the second case
	if x > 1 && y <= 1 {
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune(10)
	}
	// the third case
	if x <= 1 && y <= 1 {
		z01.PrintRune('o')
		z01.PrintRune(10)
	}
	// the 4th case
	if x <= 1 && y > 1 {

		for i := 0; i < y; i++ {
			if i == 0 || i == y-1 {
				z01.PrintRune('o')
				z01.PrintRune(10)
			} else {
				z01.PrintRune('|')
				z01.PrintRune(10)
			}
		}
	}
}
func main() {
	Raid1a(5, 3)
}
