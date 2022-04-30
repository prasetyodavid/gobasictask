package main

import (
	"fmt"
	"math"
)

func main() {
	//soalpertama()
	//soalkedua()
	//soalketiga()
	//soalkeempat()
	//soalkelima()
	//soalkeenam()
	//soalketujuh()
	//soalkedelapan()
}

func soalpertama() {
	//menghitung luas permukaan tabung
	const phi float64 = 3.14
	t := 20.0 //tingggi
	r := 4.0  //jari-jari
	fmt.Scanf("%g", &t)
	fmt.Scanf("%g", &r)
	var lp float64 = 0.0
	lp = (2 * phi * math.Pow(r, 2)) + (2 * phi * r * t)
	fmt.Println(lp)
}

func soalkedua() {
	//percabangan
	var studentScore int = 80
	fmt.Scanf("%d", &studentScore)
	switch {
	case studentScore >= 80 && studentScore <= 100:
		fmt.Println("Nilai A")
	case studentScore >= 65 && studentScore <= 79:
		fmt.Println("Nilai B")
	case studentScore >= 50 && studentScore <= 64:
		fmt.Println("Nilai C")
	case studentScore >= 35 && studentScore <= 49:
		fmt.Println("Nilai D")
	case studentScore >= 0 && studentScore <= 34:
		fmt.Println("Nilai D")
	default:
		fmt.Println("Invalid Value")
	}

}

func soalketiga() {
	//faktor bilangan
	var val int = 20
	fmt.Scanf("%d", &val)
	for i := 1; i <= val; i++ {
		if (val % i) == 0 {
			fmt.Println(i)
		}
	}
}

func isPrimeNumber(number int) bool {
	var res bool
	var count int = 0
	for i := 2; i <= number; i++ {
		if (number % i) == 0 {
			count = count + 1
		}
	}
	if count == 1 {
		//bilangan prima
		res = true
	} else {
		res = false
	}
	return res

}

func soalkeempat() {
	//bilangan prinma
	fmt.Println(isPrimeNumber(11))
	fmt.Println(isPrimeNumber(13))
	fmt.Println(isPrimeNumber(17))
	fmt.Println(isPrimeNumber(20))
	fmt.Println(isPrimeNumber(35))
	fmt.Println(isPrimeNumber(53))
	fmt.Println(isPrimeNumber(59))
	fmt.Println(isPrimeNumber(98))
	fmt.Println(isPrimeNumber(97))
}

func isPalindrome(word string) bool {
	var res bool
	var check1 string = ""
	var check2 string = ""
	for i := 0; i < len(word); i++ {
		check1 = check1 + string(word[i])
	}
	for i := len(word) - 1; i >= 0; i-- {
		check2 = check2 + string(word[i])
	}
	if check1 == check2 {
		println(check1 + " vs " + check2)
		res = true
	} else {
		println(check1 + " vs " + check2)
		res = false
	}
	return res
}

func soalkelima() {
	//palindrome
	fmt.Println(isPalindrome("civic"))
	fmt.Println(isPalindrome("katak"))
	fmt.Println(isPalindrome("kasur rusak"))
	fmt.Println(isPalindrome("naruto"))
	fmt.Println(isPalindrome("abuba"))
}

func exponentiation(num1 float64, num2 float64) float64 {
	return float64(math.Pow(num1, num2))
}

func soalkeenam() {
	//exponentiation
	fmt.Println(exponentiation(2, 3))
	fmt.Println(exponentiation(7, 2))
}

func soalketujuh() {
	//pohon natal
	max := 15
	center := max
	star := 1
	for i := 0; i < max; i++ {
		for x := 0; x < center; x++ {
			fmt.Printf(" ")
		}
		for x := 0; x < star; x++ {
			fmt.Printf("* ")
		}
		center = center - 1
		star = star + 1
		fmt.Printf("\n")
	}
}

func soalkedelapan() {
	//table perkalian
	for i := 1; i <= 9; i++ {
		fmt.Print(i)
		fmt.Print("  ")
		for x := 2; x <= 9; x++ {
			res := (x * i)
			fmt.Print(res)
			if res >= 10 {
				fmt.Print(" ")
			} else {
				fmt.Print("  ")
			}

		}
		fmt.Print("\n")
	}
}
