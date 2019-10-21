package main

import "github.com/01-edu/z01"

func Raid1d(x, y int) {
	if x > 1 && y > 1 {
		z01.PrintRune('A')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune(10)
		for i:= 0; i < y-2; i++{
			z01.PrintRune('B')
			for i := 0; i < x-2; i++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('B')
			z01.PrintRune(10)
		}
		z01.PrintRune('A')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune(10)
	}
	if x > 1 && y <= 1 {
		z01.PrintRune('A')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune(10)
	}
	if x <= 1 && y <= 1 {
		z01.PrintRune('A')
		z01.PrintRune(10)
	}
	if x <= 1 && y > 1 {
		z01.PrintRune('A')
		z01.PrintRune(10)
		for i := 0; i < y-2; i++ {
			z01.PrintRune('B')
			z01.PrintRune(10)
		}
		z01.PrintRune('A')
		z01.PrintRune(10)
	}
}
func main() {
	Raid1d(5, 4)
}
