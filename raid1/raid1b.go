package main

import "github.com/01-edu/z01"

func Raid1b(x, y int) {
	// the fist case
	if x > 1 && y > 1 {
		// 1-строка
		z01.PrintRune('/')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('*')
		}
		z01.PrintRune('\\')
		z01.PrintRune(10)
		// 2-строка
		for i := 0; i < y-2; i++ {
			z01.PrintRune('*')
			for i := 0; i < x-2; i++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('*')
			z01.PrintRune(10)
		}
		// 3-строка
		z01.PrintRune('\\')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('*')
		}
		z01.PrintRune('/')
		z01.PrintRune(10)
	}
	// the second case
	if x > 1 && y <= 1 {
		z01.PrintRune('/')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('*')
		}
		z01.PrintRune('\\')
		z01.PrintRune(10)
	}
	// the 3th case
	if x <= 1 && y <= 1 {
		z01.PrintRune('/')
		z01.PrintRune(10)
	}
	//the 4th case
	if x <= 1 && y > 1 {
		z01.PrintRune('/')
		z01.PrintRune(10)
		for i := 0; i < y-2; i++ {
			z01.PrintRune('*')
			z01.PrintRune(10)
		}
		z01.PrintRune('\\')
		z01.PrintRune(10)
	}
}
func main() {
	Raid1b(9, 1)
}
