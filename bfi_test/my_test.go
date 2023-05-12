package stringutils

import (
    "testing"
    "bytes"
    "unicode"
)

func TestSwapCase(t *testing.T) {
    input,expected := "Hello, World", "hELLO, wORLD"
    result := SwapCase(input)

    if result != expected{
        t.Errorf("SwapCase(%q) == %q, expected %q",input,result,expected)
    }
}

func TestReverse(t *testing.T) {
    input,expected := "Hello, World", "dlroW ,olleH"
    result := Reverse(input)

    if result != expected{
        t.Errorf("SwapCase(%q) == %q, expected %q",input,result,expected)
    }
}

func SwapCase(str string) string {
    buf := &bytes.Buffer{}
    for _,r := range str{
        if unicode.IsUpper(r){
            buf.WriteRune(unicode.ToLower(r))
        }else{
            buf.WriteRune(unicode.ToUpper(r))
        }
    }

    return buf.String()
}

func Reverse(s string) string{
    r:= []rune(s)
    for i,j := 0, len(r)-1; i < len(r)/2; i,j=i+1, j-1{
        r[i],r[j] = r[j],r[i]
    }

    return string(r)
}
