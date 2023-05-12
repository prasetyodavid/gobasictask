/*

Salah satu alasan kenapa OOP pada Golang lebih efisien, karena untuk mendeklarasikan private atau public method atau properti, tidak menggunakan keyword private/public. Tetapi menggunakan awalan kapital untuk public dan huruf kecil untuk private, contohnya

*/

package user

import "fmt"

type user struct {
   name string
   Age  int
}

func (u user) getName() string {
   return u.name
}

func (u user) GetUserAge() string {
   return fmt.Sprintf("%s is %d years old", u.getName(), u.Age)
}

func NewUser(name string, age int) user {
   return user{name: name, Age: age}
}
