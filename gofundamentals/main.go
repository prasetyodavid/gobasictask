package main

import (
	"bufio"
	"fmt"
	"gofundamentals/helper"
	"os"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("Simple Go examples : ")
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	opt, err := buf.ReadBytes('\n')
	optstr := strings.TrimSpace(string(opt))
	if err != nil {
		fmt.Println(err)
	} else {
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
		default:
			fmt.Println(optstr)
		}

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
	var arr [10]string
	arr[2] = "array ke-dua"
	arr[4] = "array ke-empat"
	helper.Var_dump(arr)
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
