package helper

import (
	"fmt"
	"strings"
)

func Var_dump(expression interface{}) {
	/*Global method for var_dump*/
	fmt.Printf("%#v", expression)
}

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}
