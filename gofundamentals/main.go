package main

import (
	"fmt"
	"gobasictask/gofundamentals/helper"
	"math"
	"reflect"
	"strconv"
)

func init() {
	fmt.Println("Simple Go examples : ")
}

func main() {
	var optstr string
	fmt.Scanln(&optstr)
	switch optstr {
	case "1":
		printSomething()
	case "2", "3":
		printWithParam("Hello Go.")
	case "4":
		setVariables()
	case "5":
		operatorLogic()
	case "6":
		sampleArray()
	case "7":
		sampleForLoop()
	case "8":
		sampleDefer()
	case "9":
		sampleCondition()
	case "10":
		sampleFalltrough()
	case "quiz1":
		quiz1()
	case "11":
		stringLoop()
	case "12":
		multipleReturnFunc()
	default:
		fmt.Println(optstr)
	}

}

func printSomething() {
	fmt.Println("This is Go lang !")
}

func printWithParam(param string) {
	fmt.Println(param)
}

func setVariables() {
	var myNumber int = 99
	const myConsNumber32 float32 = 888762.362323
	const myConsNumber64 float64 = 735236.345343
	myInterfaceVar1 := 345
	myInterfaceVar2 := "auto defined as string"
	fmt.Println(myNumber)
	fmt.Println(myConsNumber32)
	fmt.Println(myConsNumber64)
	myNumber = myNumber + 100
	fmt.Println(myNumber)
	fmt.Println(myInterfaceVar1)
	fmt.Println(myInterfaceVar2)
}

func operatorLogic() {
	const a int = 5
	const b int = 8
	const c string = "5"
	const d int = 8
	if a == b {
		fmt.Print("same value")
	}
	if a != b {
		fmt.Print("not same value")
	}
	if c == "5" && b == d {
		fmt.Print("both are same value")
	}
	if a >= b {
		fmt.Print("a is more than / same with b")
	} else {
		fmt.Print("a is less than b")
	}
}

func sampleArray() {
	/* Arrays */
	var arr [11]string
	x := [5]int{1, 3, 5, 6, 3}
	y := [5]int{7, 4, 3}
	arr[2] = "array ke-dua"
	arr[4] = "array ke-empat"
	fmt.Println(x)
	fmt.Println(y)
	//reflect , check type ov vars
	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(reflect.ValueOf(arr).Kind())
	helper.Var_dump(arr)
	//loop
	for idx, el := range x {
		fmt.Println(idx, "=>", el)
	}

	idx := 0
	for range y {
		fmt.Println(idx, "=>", y[idx])
		idx++
	}

	/* Slices */
	var even_num []int
	var odd_num = []int{1, 3, 5, 7, 9}
	numbers := []int{1, 2, 3, 4, 5}
	var primes = make([]int, 5, 10) // len(b)=5, cap(b)=10
	fmt.Println("Slices : ")
	fmt.Println(even_num)
	fmt.Println(odd_num)
	fmt.Println(numbers)
	fmt.Println(primes)

	/*append & copy*/
	fmt.Println("append & copy : ")
	var colors = []string{"red", "green", "yellow"}
	colors = append(colors, "purple")
	colors = append(colors, "gray")
	colors = append(colors, "black")
	colors = append(colors, "cyan")
	copied_colors := make([]string, 10)
	copied_colors_sl := colors
	copy(copied_colors, colors)
	fmt.Println(copied_colors)
	fmt.Println(copied_colors_sl)

	/*Map*/
	var mymap = map[string]int{"cireng": 1000, "tahu": 3000}
	fmt.Println(mymap)
	var price = make(map[string]int)
	price["bakwan"] = 2000
	fmt.Println(price)

	fmt.Println("3 dimensional array : ")
	const m, n, p = 2, 2, 3
	var d3 [m][n][p]float64
	d3[1][0][2] = 9.55
	fmt.Println(d3)

	fmt.Println("3 dimensional slice : ")

	var xs, ys = 5, 6
	var zs = 7
	var world = make([][][]int, xs) // x axis
	for x := 0; x < xs; x++ {
		world[x] = make([][]int, ys) // y axis
		for y := 0; y < ys; y++ {
			world[x][y] = make([]int, zs) // z axis
			for z := 0; z < zs; z++ {
				world[x][y][z] = (x+1)*100 + (y+1)*10 + (z+1)*1
			}
		}
	}
	fmt.Println(world)

}

func sampleForLoop() {

	for i := 0; i < 10; i++ {
		fmt.Println("Loop ke - " + strconv.Itoa(i))
	}

	fmt.Println("Reversed Pyramid : ")

	var max int = 10
	for i := 0; i < 10; i++ {
		for n := 0; n < max; n++ {
			fmt.Print("*")
		}
		fmt.Println("")
		max = max - 1
	}

}

func sampleDefer() {
	defer fmt.Println("Run after function finished")
	fmt.Println("run line 1")
	fmt.Println("run line 2")
	fmt.Println("run line 3")
	fmt.Println("finish")
}

func sampleCondition() {
	var a int = 1
	var b int = 5
	if a+b == 6 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}

func sampleFalltrough() {
	i := 25
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i < 30:
		fmt.Println("i is less than 30")
		fallthrough
	case i < 40:
		fmt.Println("i is also less than 40")
	}
}

func quiz1() {
	//quiz day 1
	e := []int{1, 2, 3}
	e = append(e, 4)
	fmt.Println(e, len(e))
	f := make(map[string]int)
	f["one"] = 1
	f["two"] = 2
	fmt.Println(f, len(f), f["one"], f["three"])
}

func stringLoop() {
	sentence := "Hello"

	for i := 0; i < len(sentence); i++ {
		fmt.Printf(string(sentence[i]) + "-")
	}
	for pos, char := range sentence {
		fmt.Printf("character %c start at byte position %d \n", char, pos)
	}
}

func calculateCircle(diameter float64) (float64, float64) {
	var keliling = math.Pi * math.Pow(diameter, 2)
	var luas = math.Pi * diameter
	return keliling, luas
}

func multipleReturnFunc() {
	diameter := 5.0
	keliling, luas := calculateCircle(diameter)
	fmt.Println(keliling, luas)
}
