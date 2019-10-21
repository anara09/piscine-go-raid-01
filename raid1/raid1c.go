package main

import "github.com/01-edu/z01"

func Raid1c(x, y int) {
	// the first case
	if x > 1 && y > 1 {
		// 1th line
		z01.PrintRune('A')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('A')
		z01.PrintRune(10)
		//2nd line
		for i := 0; i < y-2; i++ {
			z01.PrintRune('B')
			for i := 0; i < x-2; i++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('B')
			z01.PrintRune(10)
		}
		//3th line
		z01.PrintRune('C')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune(10)
	}
	// the 2nd case
	if x > 1 && y <= 1 {
		z01.PrintRune('A')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('A')
		z01.PrintRune(10)
	}
	// the third case
	if x <= 1 && y <= 1 {
		z01.PrintRune('A')
		z01.PrintRune(10)
	}
	//4th case
	if x <= 1 && y > 1 {
		z01.PrintRune('A')
		z01.PrintRune(10)
		for i := 0; i < y-2; i++ {
			z01.PrintRune('B')
			z01.PrintRune(10)
		}
		z01.PrintRune('C')
		z01.PrintRune(10)
	}
}
func main() {
	Raid1c(5, 4)
}
