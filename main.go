package main

import (
	"fmt"
	"math"
)

func main() {
	soalpertama()
	soalkedua()
	soalketiga()
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
